package utils

import (
	"net/http"
	"sync"
)

type Response interface {
	StatusCode() int

	Response() interface{}

	Headers() map[string]string

	SetHeader(key, val string) Response
}

type response struct {
	data interface{}

	statusCode int

	mu      sync.Mutex
	headers map[string]string
}

func (r *response) StatusCode() int {
	return r.statusCode
}

func (r *response) Response() interface{} {
	return r.data
}

func (r *response) Headers() map[string]string {
	return r.headers
}

func (r *response) SetHeader(key, val string) Response {
	r.mu.Lock()
	r.headers[key] = val
	r.mu.Unlock()
	return r
}

func OK(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusOK, headers: map[string]string{}}
}

func Created(data interface{}) Response {
	return &response{data: data, statusCode: http.StatusCreated, headers: map[string]string{}}
}
