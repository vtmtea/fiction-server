package chapter

import "vtmtea.com/fiction/model"

func GetBookChapterCount(bookId int32) int64 {
	var count int64
	model.DB.Self.Where("book_id = ?", bookId).Count(&count)

	return count
}
