package author

import (
	"github.com/sirupsen/logrus"
	"vtmtea.com/fiction/model"
)

func Create(name string) model.Author {
	newAuthor := model.Author{
		Name: name,
	}
	author := model.Author{}
	model.DB.Self.Where("name = ?", name).Find(&author)

	if author.ID > 0 {
		logrus.Println("Author already exists")
		return author
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
