package tracker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/xoebus/go-tracker/resources"
)

type ProjectClient struct {
	id     int
	client Client
}

func (p ProjectClient) Stories() (stories []resources.Story, err error) {
	query := url.Values{}
	query.Set("date_format", "millis")
	query.Set("with_state", "finished")
	request, err := p.createRequest("GET", "/stories?"+query.Encode())
	if err != nil {
		return stories, err
	}

	err = p.client.do(request, &stories)
	return stories, err
}

func (p ProjectClient) DeliverStory(storyId int) error {
	url := fmt.Sprintf("/stories/%d", storyId)
	request, err := p.createRequest("PUT", url)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Body = ioutil.NopCloser(strings.NewReader(`{"current_state":"delivered"}`))

	return p.client.do(request, nil)
}

func (p ProjectClient) createRequest(method string, path string) (*http.Request, error) {
	projectPath := fmt.Sprintf("/projects/%d%s", p.id, path)
	request, err := p.client.createRequest(method, projectPath)
	if err != nil {
		return nil, err
	}

	return request, nil
}
