package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "hyj123"
	protocol = "tcp"
	address  = "localhost:3306"
	dbname   = "gotestdb"
)

// Connector handles the db handler
var Connector *sql.DB

// InitDB connection
func InitDB() error {
	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  protocol,
		Addr:                 address,
		DBName:               dbname,
		AllowNativePasswords: true,
	}

	var err error
	Connector, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	// 链接数据库
	err = Connector.Ping()
	if err != nil {
		return err
	}

	return nil
}

func IsNoRow(err error) bool {
	if err == sql.ErrNoRows {
		return true
	}

	return false
}
