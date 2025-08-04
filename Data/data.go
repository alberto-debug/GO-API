package data

import models "github.com/alberto-debug/GO-API/Models"

var books = []models.Book{

	{ID: "1", Title: "Programming 1", Author: "Alberto junior", Year: 2025},
	{ID: "2", Title: "Data Science", Author: "Araujo", Year: 2020},
}

func GetAllBooks() []models.Book {
	return books
}

func GetBookByID(id string) (models.Book, error) {
	for _, book := range books {

		if book.ID == id {

			return book, nil

		}

	}
}
