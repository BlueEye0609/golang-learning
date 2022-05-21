package main

import (
	"GoDatabase/models"
	"fmt"
	"log"
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
			log.Fatal("Query failed. error:", err)
		}

		if bookInfo != nil {
			fmt.Printf("id: %d, title: %s, author: %s, price: %.2f\n", bookInfo.ID, bookInfo.Title, bookInfo.Author, bookInfo.Price)
		}
	}
}
