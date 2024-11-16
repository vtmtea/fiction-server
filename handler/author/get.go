package author

import (
	"vtmtea.com/fiction/model"
)

func Get(name string) model.Author {
	author := model.Author{}
	model.DB.Self.Where("name = ?", name).Find(&author)

	if author.ID > 0 {
		return author
	}

	return Create(name)
}
