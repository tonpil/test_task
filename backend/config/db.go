package config

import (
	"fmt"
	"os"
	"time"
)

type DBConfig struct {
	DBUser         string
	DBPassword     string
	DBHost         string
	DBPort         string
	DBName         string
	DBSSLMode      string
	MaxConnections int
	ConnLifetime   time.Duration
	DBURL          string
}

func NewDBConfig() *DBConfig {
	host := os.Getenv("POSTGRES_DB_HOST")
	user := os.Getenv("POSTGRES_DB_USER")
	password := os.Getenv("POSTGRES_DB_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")
	port := os.Getenv("POSTGRES_DB_PORT")
	sslMode := os.Getenv("POSTGRES_DB_SSL_MODE")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslMode)

	return &DBConfig{
		DBUser:         user,
		DBPassword:     password,
		DBHost:         host,
		DBPort:         port,
		DBName:         dbname,
		DBURL:          dbURL,
		MaxConnections: 10,
		ConnLifetime:   30 * time.Minute,
	}
}
