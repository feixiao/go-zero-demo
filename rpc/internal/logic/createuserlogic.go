package logic

import (
	"context"
	"github.com/feixiao/go-zero-demo/rpc/internal/svc"
	"github.com/feixiao/go-zero-demo/rpc/model"
	"github.com/feixiao/go-zero-demo/rpc/rpc"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *rpc.CreateUserRequest) (*rpc.CreateUserResponse, error) {
	// todo: add your logic here and delete this line

	hash, err := Encrypt(in.Password)
	if err != nil {
		return nil, err
	}

	// 插入数据库
	res, err := l.svcCtx.UserModel.Insert(model.TbUsers{
		Username: in.Username,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}

	// 返回用户ID
	uid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	logx.Infof("mobile:%s user_id:%d", in.Username, uid)

	return &rpc.CreateUserResponse{
		UserID: strconv.Itoa(int(uid)),
	}, nil
}
