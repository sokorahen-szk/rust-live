package postgresql

import (
	"fmt"

	cfg "github.com/sokorahen-szk/rust-live/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: c.PostgreSql.SkipDefaultTransaction,
	})
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

func (ps *PostgreSql) Get(value interface{}, query interface{}, args ...interface{}) error {
	resultTx := ps.db.Debug().Where(query, args...).First(value)
	if resultTx.Error != nil {
		return resultTx.Error
	}

	return nil
}

func (ps *PostgreSql) Update(value interface{}, updateValues interface{}) error {
	resultTx := ps.db.Model(value).Clauses(clause.Returning{}).Updates(updateValues)
	if resultTx.Error != nil {
		return resultTx.Error
	}

	return nil
}

func (ps *PostgreSql) List(value interface{}, query interface{}, args ...interface{}) error {
	var resultTx *gorm.DB
	if len(query.(string)) == 0 {
		resultTx = ps.db.Debug().Find(value)
	} else {
		resultTx = ps.db.Debug().Where(query, args...).Find(value)
	}

	if resultTx.Error != nil {
		return resultTx.Error
	}

	return nil
}

func (ps *PostgreSql) Truncate(tables []string) error {
	for _, table := range tables {
		resultTx := ps.db.Debug().Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if resultTx.Error != nil {
			return resultTx.Error
		}
	}

	return nil
}
