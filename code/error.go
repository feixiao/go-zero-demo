package code

type (
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func NewError(code int, msg string) *Error {
	err := &Error{Code: code, Msg: msg}
	return err
}

func (e *Error) Error() string {
	return e.Msg
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// Data 为了支持http自定义错误信息 https://go-zero.dev/cn/error-handle.html
func (e *Error) Data() *ErrorResponse {
	return &ErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
