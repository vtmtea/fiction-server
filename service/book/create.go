package book

import "vtmtea.com/fiction/model"

func Create(book model.Book) model.Book {
	model.DB.Self.Omit("deletedAt").Create(&book)

	return book
}
