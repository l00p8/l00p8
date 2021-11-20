package l00p8

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (Response, Error)
}
