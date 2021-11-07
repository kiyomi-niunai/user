package svc

import (
	"github.com/kiyomi-niunai/user/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)
import "github.com/kiyomi-niunai/user/model"

type ServiceContext struct {
	Config config.Config
	UserModel   model.UsersModel

}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	user := model.NewUsersModel(conn, c.CacheRedis)
	return &ServiceContext{
		Config: c,
		UserModel: user,
	}
}
