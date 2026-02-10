package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string     `yaml:"env" env-default:"development"`
	GRPC     GRPCServer `yaml:"grpc"`
	Database Database   `yaml:"database"`
	JWT      JWTToken
}

type GRPCServer struct {
	Port int `yaml:"port"`
}

type JWTToken struct {
	Secret string `yaml:"secret"`
	TTL    time.Duration `yaml:"ttl"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port"`
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
func (db Database) DSN() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
        db.User, db.Password, db.Host, db.Port, db.DBname)
}