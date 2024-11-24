package author

import (
	"vtmtea.com/fiction/model"
)

func Create(name string) model.Author {
	newAuthor := model.Author{
		Name: name,
	}
	model.DB.Self.Omit("deletedAt").Create(&newAuthor)
	return newAuthor
}

func CreateMultiple(names []string) {
	authors := make([]model.Author, len(names))
	for _, name := range names {
		authors = append(authors, model.Author{Name: name})
	}
	model.DB.Self.Omit("deletedAt").Create(&authors)
}
