package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "salt_tugas"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
