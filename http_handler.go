package l00p8

import (
	"net/http"
)

type Handler func(r *http.Request) Response

type Middleware func(handler Handler) Handler

type OutputMiddleware func(handler Handler) http.HandlerFunc
