package book

import "vtmtea.com/fiction/model"

func GetBySourceUrl(sourceUrl string) model.Book {
	bookModel := model.Book{
		SourceURL: sourceUrl,
	}

	model.DB.Self.Where("source_url = ?", sourceUrl).Find(&bookModel)

	return bookModel
}
