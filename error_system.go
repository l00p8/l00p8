package utils

import (
	"net/http"
	"strconv"
)

type ErrorSystem interface {
	BadRequest(code int, message ...string) Error

	InternalServerError(code int, message ...string) Error

	NotFound(code int, message ...string) Error

	Forbidden(code int, message ...string) Error

	Unauthorized(code int, message ...string) Error
}

type httpErrorSystem struct {
	system string
}

func NewErrorSystem(system string) ErrorSystem {
	return &httpErrorSystem{system}
}

func (sys httpErrorSystem) NewError(code int, status int, messages ...string) Error {
	err := &HttpError{}
	message := ""
	devMessage := ""
	if len(messages) == 2 {
		message = messages[0]
		devMessage = messages[1]
	} else if len(messages) == 1 {
		message = messages[0]
	}

	err.Code = code
	err.Status = status
	err.System = sys.system
	err.statusCode = err.Status
	err.Message = message
	err.DevMessage = devMessage
	err.MoreInfo = "errors/" + err.System + "#" + strconv.Itoa(err.Status) + "." + strconv.Itoa(err.Code)

	return err
}

func (sys httpErrorSystem) BadRequest(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusBadRequest, messages...)
}

func (sys httpErrorSystem) InternalServerError(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusInternalServerError, messages...)
}

func (sys httpErrorSystem) NotFound(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusNotFound, messages...)
}

func (sys httpErrorSystem) Forbidden(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusForbidden, messages...)
}

func (sys httpErrorSystem) Unauthorized(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusUnauthorized, messages...)
}
