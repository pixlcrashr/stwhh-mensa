package db

import (
	"errors"
	"github.com/pixlcrashr/stwhh-mensa/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func FromConfig(c config.Database) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch c.Type {
	case config.SqliteDatabaseType:
		dialector = sqlite.Open(c.DSN)
		break
	case config.PostgresqlDatabaseType:
		dialector = postgres.Open(c.DSN)
		break
	default:
		return nil, errors.New("invalid database type specified")
	}

	return gorm.Open(
		dialector,
		&gorm.Config{
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		},
	)
}
