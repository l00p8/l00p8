package l00p8

type Error interface {
	Response

	Error() string
}

type HttpError struct {
	response
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	System  string `json:"system"`
	Message string `json:"message"`
	DevMessage string `json:"developer_message,omitempty"`
	MoreInfo string `json:"more_info,omitempty"`
}

func (err *HttpError) Error() string {
	return err.Message
}

func (err *HttpError) Response() interface{} {
	return err
}
