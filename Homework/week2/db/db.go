package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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

// InitDBConn init db connection
func init() {
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
		errors.Wrap(err, "OpenSql failed")
	}

	// 链接数据库
	err = Connector.Ping()
	if err != nil {
		errors.Wrap(err, "PingSql failed")
	}

}
