package initialize

import (
	"demo/common/utils/consul"
	"demo/user/global"
	"encoding/json"
	"fmt"
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
	fmt.Printf("从本地文件中读取配置中心地址：consul-addr=%s:%d\n",
		global.ConsulConfig.Host, global.ConsulConfig.Port)

	// 从配置中心获取配置
	consulClient := consul.NewClient("consul", 8500)
	kv := consulClient.KV()
	pair, _, err := kv.Get(global.ConsulConfig.Key, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	if err := json.Unmarshal(pair.Value, &global.ServerConfig); err != nil {
		fmt.Println("配置读取失败")
		return
	}
	fmt.Printf("%#v\n", global.ServerConfig)
}
