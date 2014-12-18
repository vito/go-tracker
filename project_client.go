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

func (p ProjectClient) Stories() ([]resources.Story, error) {
	var stories []resources.Story

	query := url.Values{}
	query.Set("date_format", "millis")
	query.Set("with_state", "finished")
	request, err := p.createRequest("GET", "/stories?"+query.Encode())
	if err != nil {
		return stories, err
	}

	response, err := p.client.sendRequest(request)
	if err != nil {
		return stories, err
	}

	if err := p.client.decodeResponse(response, &stories); err != nil {
		return stories, err
	}

	return stories, nil
}

func (p ProjectClient) DeliverStory(storyId int) error {
	url := fmt.Sprintf("/stories/%d", storyId)
	request, err := p.createRequest("PUT", url)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Body = ioutil.NopCloser(strings.NewReader(`{"current_state":"delivered"}`))

	_, err = p.client.sendRequest(request)
	if err != nil {
		return err
	}

	return nil
}

func (p ProjectClient) createRequest(method string, path string) (*http.Request, error) {
	projectPath := fmt.Sprintf("/projects/%d%s", p.id, path)
	request, err := p.client.createRequest(method, projectPath)
	if err != nil {
		return nil, err
	}

	return request, nil
}
