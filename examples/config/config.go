package config

import "lark/pkg/conf"

type Config struct {
	Name       string      `yaml:"name"`
	ServerID   int         `yaml:"server_id"`
	Log        string      `yaml:"log"`
	GrpcServer *conf.Grpc  `yaml:"grpc_server"`
	Etcd       *conf.Etcd  `yaml:"etcd"`
	Mysql      *conf.Mysql `yaml:"mysql"`
	Redis      *conf.Redis `yaml:"redis"`
}

func GetConfig() *Config {
	return &Config{
		Name:       "",
		ServerID:   0,
		Log:        "",
		GrpcServer: nil,
		Etcd:       nil,
		Mysql:      nil,
		Redis:      nil,
	}
}
