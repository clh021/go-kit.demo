package conf

import "os"

var (
	HttpHost string
	HttpPort int
	GrpcHost string
	GrpcPort int
)

func GetEnv(name string, def string) string {
	env := os.Getenv(name)
	if env == "" {
		return def
	}
	return env
}
