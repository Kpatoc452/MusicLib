package main

import (
	"musiclib/cmd/server"
	"musiclib/database"
)

type Config struct {
	ServerCfg *server.Config
	DataBaseCfg *database.Config
}

func InitConfig() *Config{
	return &Config{
		ServerCfg: server.NewConfig(),
		DataBaseCfg: database.NewConfig(),
	}
}