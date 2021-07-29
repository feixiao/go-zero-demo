package logic

import (
	"context"
	"github.com/feixiao/go-zero-demo/rpc/model"
	"strconv"

	"github.com/feixiao/go-zero-demo/rpc/internal/svc"
	"github.com/feixiao/go-zero-demo/rpc/rpc"

	"github.com/feixiao/go-zero-demo/code"
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

	userID, err := strconv.Atoi(in.UserID)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserModel.FindOne(int64(userID))
	switch err {
	case nil:
		return &rpc.GetUserResponse{
			UserID:   strconv.FormatInt(user.Id, 10),
			Username: user.Username,
		}, nil
	case model.ErrNotFound:
		return nil, code.ErrUsernameUnRegister
	default:
		return nil, err
	}

}
