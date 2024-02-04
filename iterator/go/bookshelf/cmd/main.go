package main

import (
	"bookshelf/app/shelf"
	"fmt"
)

func main() {
	bookList := []string{
		"Book1",
		"Book2",
		"Book3",
		"Book4",
		"Book5",
		"Book6",
		"Book7",
		"Book8",
		"Book9",
		"Book10",
	}

	bookShelf := shelf.NewBookShelf(bookList)

	for bookShelf.HasNext() {
		item := bookShelf.Next()

		fmt.Println(item)
	}
}
