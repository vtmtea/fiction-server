package author

import (
	"github.com/sirupsen/logrus"
	"vtmtea.com/fiction/model"
)

func Create(name string) {
	newAuthor := model.Author{
		Name: name,
	}
	author := model.Author{}
	model.DB.Self.Where("name = ?", name).First(&author)

	if author.ID > 0 {
		logrus.Println("Author already exists")
		return
	}
	model.DB.Self.Omit("deleteAt").Create(&newAuthor)
}

func CreateMultiple(names []string) {
	authors := make([]model.Author, len(names))
	for _, name := range names {
		authors = append(authors, model.Author{Name: name})
	}
	model.DB.Self.Omit("deleteAt").Create(&authors)
}
