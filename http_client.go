package l00p8

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/l00p8/xclient"
)

func NewHttpClient(clt xclient.Client, errSys ErrorSystem) Client {
	return &httpClient{clt, errSys}
}

type httpClient struct {
	clt    xclient.Client
	errSys ErrorSystem
}

func (clt *httpClient) Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (Response, Error) {
	resp, err := clt.clt.Request(ctx, method, url, body, headers)
	if err != nil {
		//if nerr, ok := err.(*valkyrie.MultiError); ok {
		//	nerr.HasError()
		//}
		errStr := err.Error()
		if strings.Contains(errStr, "connect: connection refused") ||
			strings.Contains(errStr, ": no such host") {
			return nil, ErrWrapCtx(clt.errSys.BadGateway(13, errStr), ctx)
		} else if strings.Contains(errStr, "Client.Timeout exceeded while awaiting headers") ||
			strings.Contains(errStr, "hystrix: timeout") {
			return nil, ErrWrapCtx(clt.errSys.RequestTimeout(13, errStr), ctx)
		}
		//if nerr, ok := err.(net.Error); ok {
		//	if nerr.Timeout() {
		//		return nil, clt.errSys.RequestTimeout(13, nerr.Error())
		//	}
		//	return nil, clt.errSys.BadGateway(13, nerr.Error())
		//}
		//switch t := err.(type) {
		//case *net.OpError:
		//	if t.Op == "dial" {
		//		//println("Unknown host")
		//		return nil, clt.errSys.BadGateway(13, t.Error())
		//	} else if t.Op == "read" {
		//		//println("Connection refused")
		//		return nil, clt.errSys.BadGateway(13, t.Error())
		//	}
		//case syscall.Errno:
		//	if t == syscall.ECONNREFUSED {
		//		//println("Connection refused")
		//		return nil, clt.errSys.BadGateway(13, t.Error())
		//	}
		//}
		return nil, ErrWrapCtx(clt.errSys.InternalServerError(13, err.Error()), ctx)
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrWrapCtx(clt.errSys.InternalServerError(14, err.Error()), ctx)
	}
	if resp.StatusCode > 399 {
		verr := &HttpError{}
		verr.headers = map[string][]string{}
		err = json.Unmarshal(d, verr)
		if err != nil {
			return nil, ErrWrapCtx(clt.errSys.InternalServerError(15, err.Error()), ctx)
		}
		return nil, ErrWrapCtx(verr, ctx)
	}
	vresp := &response{
		data:       d,
		statusCode: resp.StatusCode,
		headers:    resp.Header,
	}
	return ResponseWrapCtx(vresp, ctx), nil
}
