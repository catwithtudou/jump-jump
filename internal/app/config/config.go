package config

import (
	"github.com/jwma/jump-jump/internal/app/db"
	"github.com/jwma/reborn"
	"time"
)

var config *reborn.Reborn

//getDefaultConfig 设置默认配置
func getDefaultConfig() *reborn.Config {
	d := reborn.NewConfig()
	d.SetValue("landingHosts", []string{"http://127.0.0.1:8081/"})

	return d
}

func GetConfig() *reborn.Reborn {
	return config
}

//TODO：自动初始化默认配置(采用redis存储配置实例)
func init() {
	var err error
	config, err = reborn.NewWithDefaults(db.GetRedisClient(), "j2config", getDefaultConfig())
	if err != nil {
		panic(err)
	}
	config.SetAutoReloadDuration(time.Second * 30)
	config.StartAutoReload()
}
