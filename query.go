package tracker

import (
	"fmt"
	"net/url"
)

type Query interface {
	Query() url.Values
}

type StoriesQuery struct {
	State State
	Label string

	Limit int
}

func (query StoriesQuery) Query() url.Values {
	params := url.Values{}
	params.Set("date_format", "millis")

	if query.State != "" {
		params.Set("with_state", string(query.State))
	}

	if query.Label != "" {
		params.Set("with_label", query.Label)
	}

	if query.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", query.Limit))
	}

	return params
}
