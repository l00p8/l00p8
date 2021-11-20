package l00p8

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

	PaymentRequired(code int, messages ...string) Error

	MethodNotAllowed(code int, messages ...string) Error

	NotAcceptable(code int, messages ...string) Error

	TooManyRequests(code int, messages ...string) Error

	RequestTimeout(code int, messages ...string) Error
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

func (sys httpErrorSystem) PaymentRequired(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusPaymentRequired, messages...)
}

func (sys httpErrorSystem) MethodNotAllowed(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusMethodNotAllowed, messages...)
}

func (sys httpErrorSystem) NotAcceptable(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusNotAcceptable, messages...)
}

func (sys httpErrorSystem) TooManyRequests(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusTooManyRequests, messages...)
}

func (sys httpErrorSystem) RequestTimeout(code int, messages ...string) Error {
	return sys.NewError(code, http.StatusRequestTimeout, messages...)
}
