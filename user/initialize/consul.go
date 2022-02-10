package initialize

import (
	"demo/user/conf"
	"demo/user/global"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

var ConsulClient *consulapi.Client
var ServiceID string

// ServiceRegister consul 服务注册
func ServiceRegister()  {
	cfg := consulapi.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ConsulConfig.Host, global.ConsulConfig.Port)

	var err error
	ConsulClient, err = consulapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成对应的的检查对象
	check := &consulapi.AgentServiceCheck{
		// 通过grpc，也可通过http做。
		GRPC:                           fmt.Sprintf("%s:%d", "golang", conf.GrpcPort),
		//HTTP:                           fmt.Sprintf("%s/health", conf.HttpAddr),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	// 生成注册对象
	registration := new(consulapi.AgentServiceRegistration)
	registration.Name = "user-srv"
	ServiceID = fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = ServiceID
	registration.Port = conf.GrpcPort
	registration.Tags = []string{"user", "srv"}
	registration.Address = "golang"
	registration.Check = check
	err = ConsulClient.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}