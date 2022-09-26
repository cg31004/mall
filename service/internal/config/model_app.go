package config

type AppConfig struct {
	LogConfig        LogConfig        `mapstructure:"log_config"`
	ServerConfig     ServerConfig     `mapstructure:"server_config"`
	GinConfig        GinConfig        `mapstructure:"gin_config"`
	MySQLConfig      MySQLConfig      `mapstructure:"mysql_config"`
	LocalCacheConfig LocalCacheConfig `mapstructure:"local_cache_config"`
	RedisConfig      RedisConfig      `mapstructure:"redis_config"`
}

type LogConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type ServerConfig struct {
	TimeZone int `mapstructure:"time_zone"`
}

type GinConfig struct {
	Port      string `mapstructure:"port"`
	DebugMode bool   `mapstructure:"debug_mode"`
	CorsMode  bool   `mapstructure:"cors_mode"`
}

type MySQLConfig struct {
	LogMode        bool `mapstructure:"log_mode"`
	MaxIdle        int  `mapstructure:"max_idle"`
	MaxOpen        int  `mapstructure:"max_open"`
	ConnMaxLifeSec int  `mapstructure:"conn_max_life_sec"`
}

type LocalCacheConfig struct {
	DefaultExpirationSec int `mapstructure:"default_expiration_sec"`
}

type RedisConfig struct {
	SessionExpirationSec int `mapstructure:"session_expiration_sec"`
}
