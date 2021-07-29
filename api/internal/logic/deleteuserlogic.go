package logic

import (
	"context"

	"github.com/feixiao/go-zero-demo/api/internal/svc"
	"github.com/feixiao/go-zero-demo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.GetUserRequest) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
