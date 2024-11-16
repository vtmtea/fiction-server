package category

import (
	"vtmtea.com/fiction/model"
	"vtmtea.com/fiction/util"
)

func Get(name string) model.Category {
	category := model.Category{}
	model.DB.Self.Where("name = ?", util.FormatCategoryName(name)).Find(&category)

	if category.ID > 0 {
		return category
	}

	return Create(util.FormatCategoryName(name))
}
