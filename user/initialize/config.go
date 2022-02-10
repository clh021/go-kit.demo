package initialize

import (
	"demo/user/global"
	"encoding/json"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig()  {
	// 从本地文件中读取配置中心地址
	cfgFile := "user/config-dev.yaml"
	v := viper.New()
	v.SetConfigFile(cfgFile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.ConsulConfig); err != nil {
		panic(err)
	}
	log.WithField("consul-addr",
		fmt.Sprintf("%s:%d",
			global.ConsulConfig.Host, global.ConsulConfig.Port,
		)).Info("从本地文件中读取配置中心地址")

	// 从配置中心获取配置
	cfg := consulapi.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ConsulConfig.Host, global.ConsulConfig.Port)

	consulClient, err := consulapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	kv := consulClient.KV()
	pair, _, err := kv.Get("user-srv/dev", nil)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	if err := json.Unmarshal(pair.Value, &global.ServerConfig); err != nil {
		fmt.Println("配置读取失败")
		return
	}
	fmt.Printf("%#v\n", global.ServerConfig)
}
