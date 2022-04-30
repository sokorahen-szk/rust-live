package config

import (
	"os"
	"strconv"
)

type Config struct {
	Env        string
	Port       int
	PostgreSql PostgreSql
}

type PostgreSql struct {
	Host     string
	User     string
	Password string
	DbName   string
	Tz       string
	SslMode  string
	Port     int
}

func NewConfig() *Config {
	c := &Config{}
	return c.Load()
}

func (c *Config) Load() *Config {
	return &Config{
		Env:  os.Getenv("APP_ENV"),
		Port: c.Int(os.Getenv("APP_SERVER_PORT")),
		PostgreSql: PostgreSql{
			Host:     os.Getenv("POSTGRESQL_HOST"),
			User:     os.Getenv("POSTGRESQL_USER"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			DbName:   os.Getenv("POSTGRESQL_DB_NAME"),
			Tz:       os.Getenv("POSTGRESQL_TZ"),
			SslMode:  os.Getenv("POSTGRESQL_SSL_MODE"),
			Port:     c.Int(os.Getenv("POSTGRESQL_SERVER_PORT")),
		},
	}
}

func (c *Config) Int(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return num
}
