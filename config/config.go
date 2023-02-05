package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		PG `yaml:"postgres"`
	}

	PG struct {
		URL        string `env-required:"true" yaml:"url" env:"PG_URL"`
		Migrations string `env-required:"true" yaml:"migrations" env:"PG_MIGRATIONS"`
		PoolMax    int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
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
