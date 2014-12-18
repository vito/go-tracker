package tracker

import "net/url"

type Query interface {
	Query() url.Values
}

type StoriesQuery struct {
	State State
}

func (query StoriesQuery) Query() url.Values {
	params := url.Values{}
	params.Set("date_format", "millis")

	if query.State != "" {
		params.Set("with_state", string(query.State))
	}

	return params
}
