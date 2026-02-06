package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env-default:"development"`
	HTTPServer HTTPServer `yaml:"http_server"`
	Database   Database   `yaml:"database"`
}

type HTTPServer struct {
	addr        string        `yaml:"addr" env-default:":8080"`
	timeout     time.Duration `yaml:"timeout" env-default:"15s"`
	idleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"password"`
	DBname   string `yaml:"db_name" env-default:"auth_db"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH_AUTH")	

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config dosent exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}