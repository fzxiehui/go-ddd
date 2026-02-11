package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load(configFile string) (*Config, error) {
	project_name := "ddd"
	v := viper.New()

	// 1. 默认值（兜底）
	v.SetDefault("http.port", 8080)
	v.SetDefault("db.name", "test.db")

	v.SetDefault("jwt.secret", project_name)
	v.SetDefault("jwt.expire", "24h")

	// 2. 显式配置文件（可选）
	if configFile != "" {
		v.SetConfigFile(configFile)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("read config file failed: %w", err)
		}
	}

	// 3. 环境变量（最高优先级）
	v.SetEnvPrefix(strings.ToUpper(project_name))
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	return &cfg, nil
}
