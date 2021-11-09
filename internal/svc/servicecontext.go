package svc

import (
	"fmt"
	"github.com/kiyomi-niunai/user/internal/config"
	"github.com/kiyomi-niunai/user/model"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//import "github.com/kiyomi-niunai/user/model"

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	RedisConn *redis.Redis
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	user := model.NewUsersModel(conn, c.CacheRedis)

	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	redisConn := redis.New(c.CacheRedis[0].Host)
	return &ServiceContext{
		Config:    c,
		DB:        db,
		RedisConn: redisConn,
		UserModel: user,
	}
}
