package l00p8

type Error interface {
	Response

	Error() string
}

type HttpError struct {
	response
	Status     int    `json:"status"`
	Code       int    `json:"code"`
	System     string `json:"system"`
	Message    string `json:"message"`
	DevMessage string `json:"developer_message,omitempty"`
	MoreInfo   string `json:"more_info,omitempty"`
}

func (err *HttpError) Error() string {
	return err.Message
}

func (err *HttpError) Response() interface{} {
	return err
}

func (err *HttpError) StatusCode() int {
	return err.Status
}

func (err *HttpError) Headers() map[string][]string {
	return err.headers
}

func (err *HttpError) SetHeader(key string, val []string) Response {
	err.mu.Lock()
	err.headers[key] = val
	err.mu.Unlock()
	return err
}
