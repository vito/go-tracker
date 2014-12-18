package tracker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xoebus/go-tracker"
)

var _ = Describe("Queries", func() {
	queryString := func(query tracker.Query) string {
		return query.Query().Encode()
	}

	Describe("StoriesQuery", func() {
		It("only has date_format by default", func() {
			query := tracker.StoriesQuery{}
			Ω(queryString(query)).Should(Equal("date_format=millis"))
		})

		It("can query by story state", func() {
			query := tracker.StoriesQuery{
				State: tracker.StateRejected,
			}
			Ω(queryString(query)).Should(Equal("date_format=millis&with_state=rejected"))
		})
	})
})
