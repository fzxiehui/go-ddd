package config

type Config struct {
	HTTP HTTPConfig `mapstructure:"http"`
	DB   DBConfig   `mapstructure:"db"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

type DBConfig struct {
	Name string `mapstructure:"name"`
}
