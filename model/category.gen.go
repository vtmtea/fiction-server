// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}