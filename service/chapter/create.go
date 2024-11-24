package chapter

import (
	"fmt"
	"vtmtea.com/fiction/handler/log"
	"vtmtea.com/fiction/model"
)

func CreateMultiple(chapters []*model.Chapter) {
	err := model.DB.Self.Omit("deletedAt").CreateInBatches(chapters, 300).Error
	if err != nil {
		log.Create(fmt.Sprintf("入库章节列表错误：%s", err.Error()), 2)
		return
	}
	log.Create(fmt.Sprintf("抓取章节数量: %d", cap(chapters)), 1)
}
