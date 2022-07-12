package conf

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Key string `json:"key"`
}

type ServerConfig struct {
	Name string `json:"name"`
}