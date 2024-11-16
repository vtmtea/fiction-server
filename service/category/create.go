package category

import "vtmtea.com/fiction/model"

func Create(category string) model.Category {
	newCategory := model.Category{
		Name: category,
	}
	model.DB.Self.Omit("deletedAt").Create(&newCategory)

	return newCategory
}
