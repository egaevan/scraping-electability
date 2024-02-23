package config

import (
	"database/sql"
	"strings"
)

func InitPgsqlDB() *sql.DB {
	cfg := config.DatabasePgSQL()
	configs := []string{
		"host=" + cfg.Host,
		"user=" + cfg.User,
		"password=" + cfg.Password,
		"dbname=" + cfg.Database,
		"port=" + cfg.Port,
		"TimeZone=" + cfg.TimeZone,
		"sslmode=require",
		"search_path=" + cfg.Schema,
	}
	dsn := strings.Join(configs, " ")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil
	}

	db.SetMaxOpenConns(cfg.MaxConnectionOpen)
	db.SetMaxIdleConns(cfg.MaxConnectionIdle)
	db.SetConnMaxLifetime(cfg.TimeOut)

	return db
}
