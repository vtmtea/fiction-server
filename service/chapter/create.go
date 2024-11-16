package chapter

import (
	"github.com/sirupsen/logrus"
	"vtmtea.com/fiction/model"
)

func CreateMultiple(chapters []*model.Chapter) {
	if err := model.DB.Self.Omit("deletedAt").CreateInBatches(chapters, 300).Error; err != nil {
		logrus.Errorln(err.Error())
	}
}
