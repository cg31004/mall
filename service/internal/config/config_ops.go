package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//go:generate mockery --name IOpsConfig --structname MockOpsConfig --output mock_config --outpkg mock_config --filename mock_ops.go --with-expecter

func newOpsConfig() IOpsConfig {
	obj := &OpsConfigSetup{
		v:              viper.New(),
		lastChangeTime: time.Now(),
	}

	obj.Load()

	return obj
}

type IOpsConfig interface {
	GetMySQLServiceConfig() MySQLServiceConfig
}

type OpsConfigSetup struct {
	v              *viper.Viper
	lastChangeTime time.Time

	OpsConfig OpsConfig `mapstructure:"ops_config"`
}

func (c *OpsConfigSetup) Load() {
	c.loadYaml()
}

func (c *OpsConfigSetup) GetLastChangeTime() time.Time {
	return c.lastChangeTime
}

func (c *OpsConfigSetup) loadYaml() {
	path, err := filepath.Abs("conf.d/app.yaml")
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

	c.v.OnConfigChange(func(in fsnotify.Event) {
		if err := c.v.Unmarshal(c); err != nil {
			panic(err)
		}
		c.lastChangeTime = time.Now()
	})

	c.v.WatchConfig()
}

func (c *OpsConfigSetup) GetMySQLServiceConfig() MySQLServiceConfig {
	return c.OpsConfig.MySQLServiceConfig
}
