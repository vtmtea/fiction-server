package chapter

import (
	"strings"
	"vtmtea.com/fiction/model"
)

func UpdateContent(contentCollection []string, chapterId int) {
	model.DB.Self.Model(&model.Chapter{}).Where("id = ?", chapterId).Update("content", strings.Join(contentCollection, "</n>"))
}
