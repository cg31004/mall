package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

//go:generate mockery --name IAppConfig --structname MockAppConfig --output mock_config --outpkg mock_config --filename mock_app.go --with-expecter

func newAppConfig() IAppConfig {
	obj := &AppConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Now(),
	}

	obj.Load()

	return obj
}

type IAppConfig interface {
	GetAppLogConfig() LogConfig
	GetServerConfig() ServerConfig
	GetGinConfig() GinConfig
	GetMySQLConfig() MySQLConfig
	GetLocalCacheConfig() LocalCacheConfig
	GetRedisConfig() RedisConfig
}

type AppConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	AppConfig AppConfig `mapstructure:"app_config"`
}

func (c *AppConfigSetup) Load() {
	c.loadYaml()
}

func (c *AppConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *AppConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d/config.yaml")
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	c.v.SetConfigType("yaml")
	c.v.SetConfigFile(path)
	if err := c.v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := c.v.Unmarshal(c); err != nil {
		panic(err)
	}
}

func (c *AppConfigSetup) GetAppLogConfig() LogConfig {
	return c.AppConfig.LogConfig
}

func (c *AppConfigSetup) GetServerConfig() ServerConfig {
	return c.AppConfig.ServerConfig
}

func (c *AppConfigSetup) GetGinConfig() GinConfig {
	return c.AppConfig.GinConfig
}

func (c *AppConfigSetup) GetMySQLConfig() MySQLConfig {
	return c.AppConfig.MySQLConfig
}

func (c *AppConfigSetup) GetLocalCacheConfig() LocalCacheConfig {
	return c.AppConfig.LocalCacheConfig
}

func (c *AppConfigSetup) GetRedisConfig() RedisConfig {
	return c.AppConfig.RedisConfig
}
