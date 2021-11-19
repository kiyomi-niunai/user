package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"github.com/kiyomi-niunai/user/internal/config"
	"github.com/kiyomi-niunai/user/internal/server"
	"github.com/kiyomi-niunai/user/internal/svc"
	"github.com/kiyomi-niunai/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)


	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()
	s.AddUnaryInterceptors(limitInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

//拦截验证器
func limitInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var limiter = rate.NewLimiter(rate.Limit(100), 10)
	if !limiter.Allow() {
		fmt.Println("限流了")
		return nil, nil
	}
	logx.Info("好玩")
	return handler(ctx, req)
}