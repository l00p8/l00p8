package l00p8

import (
	"encoding/json"
	"net/http"
)

func JSON(handler Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan Response)
		go func() {
			ch <- handler(r)
		}()
		var resp Response
		select {
		case <-r.Context().Done():
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			return
		case resp = <-ch:
			if resp == nil {
				return
			}
		}
		respBody := resp.Response()
		if respBody == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.StatusCode())
			return
		}
		var data []byte
		var ok bool
		var err error
		// assume that we already got json string in response
		if data, ok = respBody.([]byte); !ok {
			// we got some native data type in response so we need to serialize it
			data, err = json.Marshal(respBody)
			if err != nil {
				_, err = w.Write([]byte("{}"))
				if err != nil {
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		for k, val := range resp.Headers() {
			for _, v := range val {
				w.Header().Add(k, v)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode())
		_, err = w.Write(data)
		if err != nil {
			return
		}
	}
	return fn
}
