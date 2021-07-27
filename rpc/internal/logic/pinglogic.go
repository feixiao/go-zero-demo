package logic

import (
	"context"

	"github.com/feixiao/go-zero-demo/rpc/internal/svc"
	"github.com/feixiao/go-zero-demo/rpc/rpc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rpc.PingRequest) (*rpc.PingResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.PingResponse{}, nil
}
