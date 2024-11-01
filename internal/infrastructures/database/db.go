package database

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	MasterDSN       string
	MaxIdleConn     int
	MaxConn         int
	ConnMaxLifetime string
}

func New(cfg *DatabaseConfig) *sqlx.DB {

	db, err := sqlx.Connect("mysql", cfg.MasterDSN)
	if err != nil {
		panic(err.Error())
	}

	var conMaxLifetime time.Duration
	if cfg.ConnMaxLifetime != "" {
		duration, err := time.ParseDuration(cfg.ConnMaxLifetime)
		if err != nil {
			log.Fatal("Invalid ConnMaxLifetime value: " + err.Error())
		}

		conMaxLifetime = duration
	}

	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetMaxOpenConns(cfg.MaxConn)
	db.SetConnMaxLifetime(conMaxLifetime)

	if err := db.Ping(); err != nil {
		log.Panicln("failed to ping DB: ")
	}
	return db
}
