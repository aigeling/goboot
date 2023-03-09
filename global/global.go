package global

import "github.com/aigeling/goboot/frame/cache"

var Conf *Config = new(Config)

type Config struct {
	Debug       bool                `json:"debug"`
	DefaultDB   string              `json:"defaultDB"`
	Server      *Server             `json:"server"`
	DBConfigs   map[string]DBConfig `json:"DBConfigs"`
	RedisConfig *cache.RedisConf    `json:"redis"`
}

type Server struct {
	Port int `json:"port"`
	//jwt加密密钥
	JwtSecretKey string `json:"jwtSecretKey"`
}
