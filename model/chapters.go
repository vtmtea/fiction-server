package model

import (
	"time"
)

const TableNameChapter = "chapters"

// Chapter mapped from table <chapters>
type Chapter struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	BookID    int32     `gorm:"column:book_id" json:"book_id"`
	SourceID  int32     `gorm:"column:source_id" json:"source_id"`
	SourceURL string    `gorm:"column:source_url" json:"source_url"`
	Order     int32     `gorm:"column:order" json:"order"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	Source    Source    `gorm:"foreignKey:source_id" json:"source"`
}

// TableName Chapter's table name
func (*Chapter) TableName() string {
	return TableNameChapter
}
