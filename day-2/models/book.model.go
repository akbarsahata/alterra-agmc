package models

type BookModel struct {
	ID          int    `json:"id" faker:"oneof: 1, 2, 3"`
	Title       string `json:"title" faker:"sentence" validate:"required"`
	Author      string `json:"author" faker:"name" validate:"required"`
	Description string `json:"description" faker:"paragraph" validate:"required"`
}
