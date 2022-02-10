package conf

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Name string `json:"name"`
}