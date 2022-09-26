package config

type OpsConfig struct {
	MySQLServiceConfig MySQLServiceConfig `mapstructure:"mysql_service"`
	RedisServiceConfig RedisServiceConfig `mapstructure:"redis_service"`
	MongoServiceConfig MongoServiceConfig `mapstructure:"mongo_service"`
	FileServerConfig   FileServerConfig   `mapstructure:"file_server"`
}

type MySQLServiceConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisServiceConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	TTLSec   int    `mapstructure:"ttl_sec"`
}

type MongoServiceConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"passwd"`
	User     string `mapstructure:"user"`
	DB       string `mapstructure:"db"`
}

type FileServerConfig struct {
	Path     string `mapstructure:"path"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
