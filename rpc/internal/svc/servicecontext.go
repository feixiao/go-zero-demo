package svc

import (
	"github.com/feixiao/go-zero-demo/rpc/internal/config"
	"github.com/feixiao/go-zero-demo/rpc/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.TbUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	userModel := model.NewTbUsersModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config:    c,
		UserModel: userModel,
	}
}
