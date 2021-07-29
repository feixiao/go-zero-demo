package logic

import "github.com/feixiao/go-zero-demo/api/internal/types"

func NewOKResponse(data interface{}) *types.Response {
	return &types.Response{
		Code:    0,
		Message: "成功",
		Data:    data,
	}
}
