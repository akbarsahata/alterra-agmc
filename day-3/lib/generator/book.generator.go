package generator

import (
	"github.com/akbarsahata/alterra-agmc/day-3/models"
	"github.com/go-faker/faker/v4"
)

func GenerateBooks() []models.BookModel {
	books := []models.BookModel{}

	for i := 0; i < 3; i++ {
		fakeBook := models.BookModel{}

		faker.FakeData(&fakeBook)

		books = append(books, fakeBook)
	}

	return books
}

func GenerateBook() models.BookModel {
	book := models.BookModel{}

	faker.FakeData(&book)

	return book
}
