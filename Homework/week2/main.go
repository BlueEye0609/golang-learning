package main

import (
	"GoDatabase/models"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main() {
	queryBook()
	//queryRows()
}

func queryBook() {
	for i := 0; i < 3; i++ {
		book := &models.Book{ID: int32(i)}
		bookInfo, err := book.GetById()
		if err != nil {
			fmt.Printf("original error: %T. %v\n", errors.Cause(err), errors.Cause(err))
			fmt.Printf("stack trace:\n%+v\n", err)
			os.Exit(1)
		}

		if bookInfo != nil {
			fmt.Printf("id: %d, title: %s, author: %s, price: %.2f\n", bookInfo.ID, bookInfo.Title, bookInfo.Author, bookInfo.Price)
		}
	}
}
