package config

import (
	"os"
	"time"
)

type DatabasePGSQLConfig struct {
	User              string
	Password          string
	Database          string
	Host              string
	Port              string
	TimeZone          string
	MaxConnectionIdle int
	MaxConnectionOpen int
	Schema            string
	Debug             bool
	TimeOut           time.Duration
}

func DatabasePgSQL() DatabasePGSQLConfig {
	schema := "public"
	cfgSchema := os.Getenv("DB_SCHEMA")
	if cfgSchema != "" {
		schema = cfgSchema
	}

	return DatabasePGSQLConfig{
		User:              os.Getenv("PGSQL_USERNAME"),
		Password:          os.Getenv("PGSQL_PASS"),
		Database:          os.Getenv("PGSQL_NAME"),
		Host:              os.Getenv("PGSQL_HOST"),
		Port:              os.Getenv("PGSQL_PORT"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		MaxConnectionIdle: utils.ConvertEnvToInt("DB_MAX_CON_IDLE"),
		MaxConnectionOpen: utils.ConvertInt("DB_MAX_CON_OPEN"),
		Schema:            schema,
		Debug:             utils.ConvertBool("DEBUG"),
		TimeOut:           time.Duration(utils.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
