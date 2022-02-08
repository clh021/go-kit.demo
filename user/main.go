package main

import (
	"demo/user/initialize"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"demo/user/conf"
	"demo/user/server/grpc"
	"demo/user/server/http"
	log "github.com/sirupsen/logrus"
)

func init() {
	// flags
	flag.StringVar(&conf.HttpAddr, "http-addr", conf.GetEnv("HttpAddr", "0.0.0.0:8080"), "http服务地址")
	flag.StringVar(&conf.GrpcAddr, "grpc-addr", conf.GetEnv("GrpcAddr", "0.0.0.0:5000"), "grpc服务地址")

	// log
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2005-01-02 15:04:05",
	})

	log.WithFields(log.Fields{
		"http-addr": conf.HttpAddr,
		"grpc-addr": conf.GrpcAddr,
	}).Info("run flags:")
}

func main() {
	flag.Parse()

	errc := make(chan error)

	// http server
	{
		log.WithField("http-addr", conf.HttpAddr).Info("http server is running...")
		go http.Run(conf.HttpAddr, errc)
	}

	// grpc server
	{
		log.WithField("grpc-addr", conf.GrpcAddr).Info("grpc server is running...")
		go grpc.Run(conf.GrpcAddr, errc)
	}

	// consul 服务注册
	initialize.ServiceRegister()

	// 监听终止信号
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-quit)
	}()

	log.WithField("error", <-errc).Info("Exit")

	// 服务反注册
	if err := initialize.ConsulClient.Agent().ServiceDeregister(initialize.ServiceID); err != nil{
		log.Info("服务从consul中：注销失败")
	}
	log.Info("服务从consul中：注销成功")
}
