package l00p8

import (
	"context"
	"net/http"
	"sync"

	"github.com/go-chi/chi/middleware"
)

type Response interface {
	StatusCode() int

	Response() interface{}

	Headers() map[string][]string

	SetHeader(key string, val []string) Response
}

type response struct {
	data interface{}

	statusCode int

	mu      sync.Mutex
	headers map[string][]string
}

func (r *response) StatusCode() int {
	return r.statusCode
}

func (r *response) Response() interface{} {
	return r.data
}

func (r *response) Headers() map[string][]string {
	return r.headers
}

func (r *response) SetHeader(key string, val []string) Response {
	r.mu.Lock()
	r.headers[key] = val
	r.mu.Unlock()
	return r
}

func OK(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusOK, headers: map[string][]string{}}
}

func Created(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusCreated, headers: map[string][]string{}}
}

func Accepted(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusAccepted, headers: map[string][]string{}}
}

func NonAuthoritativeInformation(data interface{}) Response {
	return &response{data: data, statusCode: 203, headers: map[string][]string{}}
}

func NoContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusNoContent, headers: map[string][]string{}}
}

func ResetContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusResetContent, headers: map[string][]string{}}
}

func PartialContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusPartialContent, headers: map[string][]string{}}
}

func MultiStatus(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusMultiStatus, headers: map[string][]string{}}
}

func AlreadyReported(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusAlreadyReported, headers: map[string][]string{}}
}

func IMUsed(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusIMUsed, headers: map[string][]string{}}
}

func ResponseWrapCtx(resp Response, ctx context.Context) Response {
	xReqId, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		return resp
	}
	resp.SetHeader(middleware.RequestIDHeader, []string{xReqId})
	return resp
}
