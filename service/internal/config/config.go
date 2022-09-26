package config

type EnvType = string

const (
	EenTypeLocal EnvType = "local"
	EenTypeProd  EnvType = "prod"
)

func NewAppConfig() IAppConfig {
	return newAppConfig()
}

func NewOpsConfig() IOpsConfig {
	return newOpsConfig()
}
