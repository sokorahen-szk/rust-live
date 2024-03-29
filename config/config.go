package config

import (
	"os"
	"strconv"
)

type Config struct {
	Env        string
	Port       int
	PostgreSql PostgreSql
	Redis      Redis
	Batch      Batch
}

type PostgreSql struct {
	Host                   string
	User                   string
	Password               string
	DbName                 string
	Tz                     string
	SslMode                string
	Port                   int
	SkipDefaultTransaction bool
}

type Redis struct {
	Host           string
	Password       string
	Port           int
	DbNumber       int
	DefaultTtlHour int
}

type Batch struct {
	ApiTwitchClientId  string
	ApiTwtichSecretKey string

	ApiYoutubeSecretKey string
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
			Host:                   os.Getenv("POSTGRESQL_HOST"),
			User:                   os.Getenv("POSTGRESQL_USER"),
			Password:               os.Getenv("POSTGRESQL_PASSWORD"),
			DbName:                 os.Getenv("POSTGRESQL_DB_NAME"),
			Tz:                     os.Getenv("POSTGRESQL_TZ"),
			SslMode:                os.Getenv("POSTGRESQL_SSL_MODE"),
			Port:                   c.Int(os.Getenv("POSTGRESQL_SERVER_PORT")),
			SkipDefaultTransaction: c.Bool(os.Getenv("POSTGRESQL_SKIP_DEFAULT_TRANSACTION")),
		},
		Redis: Redis{
			Host:           os.Getenv("REDIS_HOST"),
			Password:       os.Getenv("REDIS_PASSWORD"),
			Port:           c.Int(os.Getenv("REDIS_SERVER_PORT")),
			DbNumber:       c.Int(os.Getenv("REDIS_DB_NUMBER")),
			DefaultTtlHour: c.Int(os.Getenv("REDIS_DEFAULT_TTL_HOUR")),
		},
		Batch: Batch{
			ApiTwitchClientId:   os.Getenv("API_TWITCH_CLIENT_ID"),
			ApiTwtichSecretKey:  os.Getenv("API_TWITCH_SECRET_KEY"),
			ApiYoutubeSecretKey: os.Getenv("API_YOUTUBE_SECRET_KEY"),
		},
	}
}

func (c *Config) Int(value string) int {
	toInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return toInt
}

func (c *Config) Bool(value string) bool {
	toBool, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return toBool
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

func (c *Config) IsDev() bool {
	return c.Env == "dev"
}

func (c *Config) IsTest() bool {
	return c.Env == "test" || c.Env == "local"
}
