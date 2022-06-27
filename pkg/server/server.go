package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"short-url/config"
	"short-url/pkg/database/mysql"
	"short-url/pkg/database/redis"
	"short-url/pkg/server/api"
	"syscall"

	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func init() {
	config.Setup()
	mysql.InitMysql()
	redis.InitRedis()
	Server = gin.New()
	addMiddleware()
	addRoute()
}

func addMiddleware() {
	Server.Use(gin.Logger())
	Server.Use(gin.Recovery())
}

func addRoute() {
	Server.GET("/s/:short", api.RedirectOriginUrl)
	v1 := Server.Group("/api/v1")
	{
		v1.GET("/short", api.GetShortUrl)
	}
}

func Run(address string) {
	server := &http.Server{
		Addr:    address,
		Handler: Server,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	select {
	case <-ctx.Done():
		cancel()
		fmt.Println("Context is cancelled")
	case <-quit:
		fmt.Println("Shutdown Server ...")
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("Server Shutdown:" + err.Error())
		}
	}
	fmt.Println("Server exiting")
}
