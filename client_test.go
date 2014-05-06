package tracker_test

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"

	"github.com/xoebus/go-tracker"
)

func Fixture(filename string) string {
	path := filepath.Join("fixtures", filename)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

var _ = Describe("Tracker Client", func() {
	var (
		server *ghttp.Server
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		tracker.DefaultURL = server.URL()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("getting information about the current user", func() {
		var statusCode int

		It("works if everything goes to plan", func() {
			statusCode = http.StatusOK
			headers := http.Header{
				"X-TrackerToken": {"api-token"},
			}

			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/services/v5/me"),
				ghttp.VerifyHeader(headers),

				ghttp.RespondWith(statusCode, Fixture("me.json")),
			))

			client := tracker.NewClient("api-token")
			me, err := client.Me()

			Ω(err).ToNot(HaveOccurred())
			Ω(me.Username).To(Equal("vader"))
		})

		It("returns an error if the response is not successful", func() {
			statusCode = http.StatusInternalServerError

			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.RespondWith(statusCode, ""),
			))

			client := tracker.NewClient("api-token")
			_, err := client.Me()
			Ω(err).To(MatchError("request failed (500)"))
		})

		It("returns a helpful error if the token is invalid", func() {
			statusCode = http.StatusUnauthorized

			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.RespondWith(statusCode, ""),
			))

			client := tracker.NewClient("api-token")
			_, err := client.Me()
			Ω(err).To(MatchError("invalid token"))
		})

		It("returns an error if the request fails", func() {
			server.Close()

			client := tracker.NewClient("api-token")
			_, err := client.Me()

			Ω(err).To(HaveOccurred())
			Ω(err.Error()).To(MatchRegexp("failed to make request"))
			server = ghttp.NewServer()
		})

		It("returns an error if the request can't be created", func() {
			tracker.DefaultURL = "aaaaa)#Q&%*(*"

			client := tracker.NewClient("api-token")
			_, err := client.Me()

			Ω(err).To(HaveOccurred())
			Ω(err.Error()).To(MatchRegexp("failed to create request"))
		})

		It("returns an error if the response JSON is broken", func() {
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.RespondWith(http.StatusOK, `{"`),
			))

			client := tracker.NewClient("api-token")
			_, err := client.Me()

			Ω(err).To(HaveOccurred())
			Ω(err.Error()).To(MatchRegexp("invalid json response"))
		})
	})
})
