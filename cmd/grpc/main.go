package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"short-url/config"
	"short-url/pkg/database/mysql"
	"short-url/pkg/database/redis"
	"short-url/pkg/grpc/service"
	"short-url/pkg/pb"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	config.Setup()
	mysql.InitMysql()
	redis.InitRedis()
	server := grpc.NewServer()
	// 监听服务关闭信号，服务平滑重启
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		<-c
		server.GracefulStop()
	}()
	pb.RegisterTinyUrlServer(server, &service.TinyUrlService{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	log.Println("rpc服务已经开启")
	err = server.Serve(listener)
	log.Println(err)
}
