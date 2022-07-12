package consul

import (
	"errors"
	"fmt"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

type Client struct {
	//Host string
	//Port int
}

type ServiceInstance struct {
	InstanceId string
	ServiceId  string
	Host       string
	Port       int
	Secure     bool
	Metadata   map[string]string
}

func NewClient(host string, port int) *consulapi.Client {
	cfg := consulapi.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", host, port)

	consulClient, err := consulapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return consulClient
}

// Registry 服务注册，实现注册和反注册两个方法
type Registry struct {
	Client *consulapi.Client
}

type RegistryClient interface {
	Register(host string, port int, name string, tags []string, id string, way string) error
	RegisterByHttp(host string, port int, name string, tags []string, id string) error
	RegisterByHttps(host string, port int, name string, tags []string, id string) error
	RegisterByGrpc(host string, port int, name string, tags []string, id string) error
	DeRegister(id string) error
	ServiceDiscovery(servername string) ([]ServiceInstance, error)
}

func NewRegistryClient(client *consulapi.Client) RegistryClient {
	return &Registry{
		Client: client,
	}
}

func (r *Registry) Register(host string, port int, name string, tags []string, id string, way string) error {
	// 生成对应的的检查对象
	check := &consulapi.AgentServiceCheck{
		// 通过grpc，也可通过http做。
		//GRPC:                           fmt.Sprintf("%s:%d", host, port),
		//HTTP:                           fmt.Sprintf("%s:%d/health", host, port),
		Timeout:                        "10s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	if strings.ToUpper(way) == "HTTP" {
		check.HTTP = fmt.Sprintf("http://%s:%d/health", host, port)
	} else if strings.ToUpper(way) == "HTTPS" {
		check.HTTP = fmt.Sprintf("https://%s:%d/health", host, port)
	} else {
		// grpc 不需要加协议名和/health路径
		check.GRPC = fmt.Sprintf("%s:%d", host, port)
		check.Notes = "Consul check service health status."
	}

	// 生成注册对象
	registration := new(consulapi.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = host
	registration.Check = check
	err := r.Client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	return nil
}

func (r *Registry) RegisterByHttp(host string, port int, name string, tags []string, id string) error {
	return r.Register(host, port, name, tags, id, "http")
}

func (r *Registry) RegisterByHttps(host string, port int, name string, tags []string, id string) error {
	return r.Register(host, port, name, tags, id, "https")
}

func (r *Registry) RegisterByGrpc(host string, port int, name string, tags []string, id string) error {
	return r.Register(host, port, name, tags, id, "grpc")
}

func (r *Registry) DeRegister(id string) error {
	err := r.Client.Agent().ServiceDeregister(id)
	return err
}

func (r *Registry) ServiceDiscovery(serviceName string) ([]ServiceInstance, error) {
	//获取所有的服务
	catalogService, _, err := r.Client.Catalog().Service(serviceName, "", nil)
	if err != nil {
		return nil, err
	}
	if len(catalogService) <= 0 {
		return nil, errors.New("no " + serviceName + " consul service")
	}
	instance := make([]ServiceInstance, len(catalogService))

	for index, server := range catalogService {
		s := ServiceInstance{
			InstanceId: server.ServiceID,
			ServiceId:  server.ServiceName,
			Host:       server.ServiceAddress,
			Port:       server.ServicePort,
			Metadata:   server.ServiceMeta,
		}
		instance[index] = s
	}

	return instance, nil
}