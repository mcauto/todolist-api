package todo

import (
	"time"

	"gorm.io/gorm"
)

// Item todo item
type Item struct {
	ID      uint64         `json:"id"      gorm:"->;primaryKey;autoIncrement"`
	Title   string         `json:"title"   gorm:"type:varchar(100);uniqueIndex"`
	Checked bool           `json:"checked" gorm:"default: false"`
	Created time.Time      `json:"created" gorm:"autoCreateTime"`
	Updated time.Time      `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"-"       gorm:"index"`
}

// TableName table name
func (Item) TableName() string {
	return "todos"
}
