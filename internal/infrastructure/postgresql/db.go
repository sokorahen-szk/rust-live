package postgresql

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
	db *gorm.DB
}

func NewPostgreSQL(c *cfg.Config) *PostgreSql {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.PostgreSql.Host,
		c.PostgreSql.User,
		c.PostgreSql.Password,
		c.PostgreSql.DbName,
		c.PostgreSql.Port,
		c.PostgreSql.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &PostgreSql{
		db: db,
	}
}

func (ps *PostgreSql) Create(value interface{}) error {
	resultTx := ps.db.Debug().Create(value)
	if resultTx.Error != nil {
		return resultTx.Error
	}

	return nil
}
