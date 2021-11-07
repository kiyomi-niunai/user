package logic

import (
	"context"
	"fmt"
	"strconv"

	//"fmt"
	//"github.com/kiyomi-niunai/user/model"
	//"github.com/tal-tech/go-zero/core/stores/cache"
	//"github.com/tal-tech/go-zero/core/stores/sqlx"

	"github.com/kiyomi-niunai/user/internal/svc"
	"github.com/kiyomi-niunai/user/user"
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

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line]
	id, _ := strconv.ParseInt(in.Id, 10, 64)
	userObj, err := l.svcCtx.UserModel.FindOne(id)
	if err != nil {
		fmt.Println("报错的是", err)
	}
	return &user.UserResponse{
		Id: strconv.Itoa(int(userObj.Id)),
		Name: userObj.Name,
	}, nil
}
