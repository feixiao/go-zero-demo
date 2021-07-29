package code

import (
	"strings"
	"sync"
)

type (
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
)

func NewError(code int, msg string) *Error {
	err := &Error{Code: code, Msg: msg}
	errors.Store(msg, err)
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

// 错误码管理

var errors sync.Map

func ToCodeError(err1 error) *Error {
	// grpc 返回的错误内容如下
	// rpc error: code = Unknown desc = 手机号已经被占用

	var err *Error
	errors.Range(func(key, value interface{}) bool {
		str := key.(string)
		if strings.Contains(err1.Error(), str) {
			err = value.(*Error)
			return false
		} else {
			return true // 继续遍历
		}
	})

	if err != nil {
		return err
	} else {
		return NewError(1000, err1.Error())
	}
}
