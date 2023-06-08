package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "SUCCESS",
		Data:    data,
	}
}

func Error(msg string) *Response {
	return &Response{
		Code:    -1,
		Message: msg,
	}
}
