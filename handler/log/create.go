package log

import "vtmtea.com/fiction/model"

func Create(logInfo string, logType int32) {
	log := model.Log{
		Type:    logType,
		Message: logInfo,
	}
	model.DB.Self.Omit("deletedAt").Create(&log)
}
