package config

type OpsConfig struct {
	MySQLServiceConfig MySQLServiceConfig `mapstructure:"mysql_service"`
}

type MySQLServiceConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
