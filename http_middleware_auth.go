package l00p8

import (
	"context"
	"github.com/l00p8/shield"
	"net/http"
	"strings"
)

type MiddlewareFactory interface {
	Auth(handler Handler) Handler
}

type auth struct {
	validator shield.Validator
	errSys    ErrorSystem
}

func NewMiddlewareFactory(validator shield.Validator, errSys ErrorSystem) MiddlewareFactory {
	return &auth{validator: validator, errSys: errSys}
}

func (auth *auth) Auth(handler Handler) Handler {
	return func(r *http.Request) Response {
		header := r.Header.Get("Authorization")
		if header == "" {
			return auth.errSys.BadRequest(10, "Bad request.")
		}
		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return auth.errSys.BadRequest(11, "Bad request.")
		}
		tkn := parts[1]
		if tkn == "" {
			return auth.errSys.BadRequest(12, "Bad request.")
		}
		claims, err := auth.validator.Valid(tkn)
		if err != nil {
			return auth.errSys.Unauthorized(13, err.Error())
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		return handler(r.WithContext(ctx))
	}
}
