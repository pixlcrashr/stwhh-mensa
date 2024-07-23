package config

type Config struct {
	Database Database `mapstructure:"db"`
}

type DatabaseType string

const (
	SqliteDatabaseType     = "sqlite"
	PostgresqlDatabaseType = "postgresql"
)

type Database struct {
	Type DatabaseType `mapstructure:"type"`
	DSN  string       `mapstructure:"dsn"`
}
