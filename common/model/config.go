package model

type Config struct {
	Port     string
	Database DbConfig
	Casbin   CasbinConfig
	JWT      JWTConfig
	Logger   LoggerConfig
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string

	MaxOpenConns int `mapstructure:"max-open-conns"`
	MaxIdleConns int `mapstructure:"max-idle-conns"`
}

type CasbinConfig struct {
	Path string
}

type JWTConfig struct {
	Secret string
	Expire int //多少小时过期
}

type LoggerConfig struct {
	Level string
}
