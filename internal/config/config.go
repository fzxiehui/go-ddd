package config

import "time"

type Config struct {
	HTTP HTTPConfig `mapstructure:"http"`
	DB   DBConfig   `mapstructure:"db"`
	JWT  JWTConfig  `mapstructure:"jwt"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

type DBConfig struct {
	Name string `mapstructure:"name"`
}

type JWTConfig struct {
	Secret string        `mapstructure:"secret"`
	Expire time.Duration `mapstructure:"expire"`
}
