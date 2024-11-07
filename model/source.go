package model

type SourceModel struct {
    BaseModel
    Name string `json:"name" gorm:"column:name;not null"`
    Url  string `json:"url" gorm:"column:url;not null"`
}

func (source *SourceModel) TableName() string {
    return "fs_source"
}
