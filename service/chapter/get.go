package chapter

import "vtmtea.com/fiction/model"

func GetBookChapterCount(bookId int32) int64 {
	var count int64
	model.DB.Self.Model(&model.Chapter{}).Where("book_id = ?", bookId).Count(&count)

	return count
}

func GetChapter(chapterId int) model.Chapter {
	var chapter model.Chapter
	model.DB.Self.Where("id = ?", chapterId).Preload("Source").Find(&chapter)
	return chapter
}
