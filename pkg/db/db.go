package db

import (
	"fmt"

	"github.com/jmoiron/sqlx" // extension for standard go database/sql lib
)

type ConfingDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(cnfg ConfingDB) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnfg.Host, cnfg.Port, cnfg.User, cnfg.Password, cnfg.Name,
	)

	db, err := sqlx.Connect("postgres", dataSource)

	return db, err
}
