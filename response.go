package l00p8

import (
	"net/http"
	"sync"
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
	return &response{data: data, statusCode: http.StatusOK, headers: nil}
}

func Created(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusCreated, headers: nil}
}

func Accepted(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusAccepted, headers: nil}
}

func NonAuthoritativeInformation(data interface{}) Response {
	return &response{data: data, statusCode: 203, headers: nil}
}

func NoContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusNoContent, headers: nil}
}

func ResetContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusResetContent, headers: nil}
}

func PartialContent(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusPartialContent, headers: nil}
}

func MultiStatus(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusMultiStatus, headers: nil}
}

func AlreadyReported(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusAlreadyReported, headers: nil}
}

func IMUsed(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusIMUsed, headers: nil}
}
