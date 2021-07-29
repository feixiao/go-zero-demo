package logic

import (
	"context"
	"github.com/feixiao/go-zero-demo/rpc/rpc"

	"github.com/feixiao/go-zero-demo/api/internal/svc"
	"github.com/feixiao/go-zero-demo/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req types.GetUserRequest) (*types.Response, error) {
	// todo: add your logic here and delete this line

	logx.Infof("GetUserInfo req:%+v", req)
	rsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &rpc.GetUserRequest{
		UserID: req.UserID,
	})

	if err != nil {

		return nil, err
	}

	data := &types.GetUserResponse{
		UserID:   rsp.UserID,
		Username: rsp.Username,
	}
	return NewOKResponse(data), nil
}
