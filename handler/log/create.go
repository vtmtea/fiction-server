package log

import (
	"github.com/sirupsen/logrus"
	"vtmtea.com/fiction/model"
)

func Create(logInfo string, logType int32) {
	log := model.Log{
		Type:    logType,
		Message: logInfo,
	}
	model.DB.Self.Omit("deletedAt").Create(&log)
}

func CreateMultiple(logs []model.Log) {
	if err := model.DB.Self.Omit("deletedAt").Create(&logs).Error; err != nil {
		logrus.Errorln(err.Error())
	}
}
