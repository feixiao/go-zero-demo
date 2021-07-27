package logic

import (
	"context"

	"github.com/feixiao/go-zero-demo/rpc/internal/svc"
	"github.com/feixiao/go-zero-demo/rpc/rpc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *rpc.GetUserRequest) (*rpc.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.GetUserResponse{}, nil
}
