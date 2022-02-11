package consul

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	errc := make(chan error)

	client := NewClient("consul", 8500)
	rc := NewRegistryClient(client)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	tags := []string{"srv"}

	// 服务注册（http方式）
	err := rc.Register("golang", 8080, "consul_test", tags, serviceId, "http")
	if err != nil {
		t.Error(err)
	}
	t.Log("服务注册成功")

	// 监听终止信号
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-quit)
	}()

	//log.WithField("error", <-errc).Info("Exit"
	// 模拟 server 运行
	time.Sleep(time.Second * 2)

	// 反注册
	if err := rc.DeRegister(serviceId); err != nil{
		t.Log("服务从consul中：注销失败")
	} else {
		t.Log("服务从consul中：注销成功")
	}
}

func TestConfigCenter(t *testing.T)  {
	client := NewClient("consul", 8500)

	// 配置中心
	kv := client.KV()
	pair, _, err := kv.Get("user-srv/dev", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("KV: %v %s\n", pair.Key, pair.Value)
	t.Log("从配置中心获取数据成功")
}
