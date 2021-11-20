package l00p8

import (
	"context"
	"encoding/json"
	httpclient "github.com/l00p8/xclient"
	"io"
	"io/ioutil"
	"net/http"
)

func NewHttpClient(clt httpclient.Client, errSys ErrorSystem) Client {
	return &httpClient{clt, errSys}
}

type httpClient struct {
	clt    httpclient.Client
	errSys ErrorSystem
}

func (clt *httpClient) Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (Response, Error) {
	resp, err := clt.clt.Request(ctx, method, url, body, headers)
	if err != nil {
		return nil, clt.errSys.InternalServerError(13, err.Error())
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, clt.errSys.InternalServerError(14, err.Error())
	}
	if resp.StatusCode > 399 {
		verr := &HttpError{}
		err = json.Unmarshal(d, verr)
		if err != nil {
			return nil, clt.errSys.InternalServerError(15, err.Error())
		}
		return nil, verr
	}
	vresp := &response{
		data: d,
		statusCode: resp.StatusCode,
		headers: resp.Header,
	}
	return vresp, nil
}
