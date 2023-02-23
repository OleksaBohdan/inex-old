package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		HTTP `yaml:"http"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		URL        string `env-required:"true" yaml:"url" env:"PG_URL"`
		Migrations string `env-required:"true" yaml:"migrations" env:"PG_MIGRATIONS"`
		PoolMax    int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"PORT"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("../config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
