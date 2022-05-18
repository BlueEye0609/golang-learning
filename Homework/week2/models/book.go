package models

import (
	"GoDatabase/db"
	"database/sql"
	"github.com/pkg/errors"
)

type Book struct {
	ID     int32
	Title  string
	Author string
	Price  float64
}

func (bk *Book) GetById() (*Book, error) {
	err := db.Connector.QueryRow("SELECT id, title, author, price FROM books WHERE id = ?", bk.ID).Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "Query failed")
		}
	}
	return bk, nil
}
