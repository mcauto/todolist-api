package models

// TodoItem 할 일의 아이템
type TodoItem struct {
	ID      uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Text    string `json:"text" gorm:"type:varchar(100);unique_index" validate:"required"`
	Checked bool   `json:"checked" gorm:"default: false"`
}
