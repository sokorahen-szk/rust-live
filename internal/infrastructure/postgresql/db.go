package postgresql

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSql struct {
	*gorm.DB
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: c.PostgreSql.SkipDefaultTransaction,
	})
	if err != nil {
		panic(err)
	}

	return &PostgreSql{
		db,
	}
}

func (ps *PostgreSql) Truncate(tables []string) error {
	for _, table := range tables {
		resultTx := ps.Debug().Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if resultTx.Error != nil {
			return resultTx.Error
		}
	}

	return nil
}
