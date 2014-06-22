package tracker

import (
	"fmt"
	"net/http"

	"github.com/xoebus/go-tracker/resources"
)

type ProjectClient struct {
	id     int
	client Client
}

func (p ProjectClient) Stories() ([]resources.Story, error) {
	var stories []resources.Story

	request, err := p.createRequest("/stories?date_format=millis&with_state=finished")
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

func (p ProjectClient) createRequest(path string) (*http.Request, error) {
	projectPath := fmt.Sprintf("/projects/%d%s", p.id, path)
	request, err := p.client.createRequest(projectPath)
	if err != nil {
		return nil, err
	}

	return request, nil
}
