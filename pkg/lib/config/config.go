package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
	URL    string `env:"API_URL" envDefault:"https://localhost:3443"`
	Env    string `env:"API_ENV" envDefault:"dev"`
	Port   string `env:"API_PORT" envDefault:"8080"`
	DBHost string `env:"API_DB_HOST" envDefault:"172.30.0.3"`
	DBName string `env:"API_DB_NAME" envDefault:"golang"`
	DBUser string `env:"API_DB_USER" envDefault:"root"`
	DBPass string `env:"API_DB_PASS" envDefault:"pass"`
	DBPort string `env:"API_DB_PORT" envDefault:"3306"`
}

// Config（環境変数）の読み取り
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("fail to parse cfg: %w", err)
	}
	return cfg, nil
}
