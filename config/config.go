package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port int
}

func NewConfig() *Config {
	c := &Config{}
	return c.Load()
}

func (c *Config) Load() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	return &Config{
		Env:  os.Getenv("APP_ENV"),
		Port: c.Int(os.Getenv("APP_SERVER_PORT")),
	}
}

func (c *Config) Int(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return num
}
