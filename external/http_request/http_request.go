package http_request

import (
	"context"
)

// simple wrapper for http request
// for simplicity we will use all as json body type

type HTTPRequest interface {
	Post(ctx context.Context, url string, header map[string]string, body interface{}, response interface{}) error
}
