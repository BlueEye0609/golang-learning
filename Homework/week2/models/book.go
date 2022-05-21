package models

import (
	"GoDatabase/db"
)

type Book struct {
	ID     int32
	Title  string
	Author string
	Price  float64
}

func (bk *Book) GetById() (*Book, error) {
	err := db.InitDB()
	if err != nil {
		return nil, err
	}
	err = db.Connector.QueryRow("SELECT id, title, author, price FROM books WHERE id = ?", bk.ID).Scan(&bk.ID, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		if ok := db.IsNoRow(err); ok {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return bk, nil
}
