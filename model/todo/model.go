package todo

import "github.com/jinzhu/gorm"

// Item todo item
type Item struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Text    string `gorm:"type:varchar(100);unique_index"`
	Checked bool   `gorm:"default: false"`
}

// FindItem find todo item
func FindItem(db *gorm.DB) (items []Item, err error) {
	err = db.Find(&items).Error
	return
}

// Create create todo item
func (i *Item) Create(db *gorm.DB) (err error) {
	err = db.Create(i).Error
	return
}

// Find find todo item
func (i *Item) Find(db *gorm.DB) (err error) {
	err = db.First(i).Error
	return
}

// Update update todo item as checked
func (i *Item) Update(db *gorm.DB) (err error) {
	err = db.Model(i).Update(i).Error
	return
}

// Delete delete todo item
func (i *Item) Delete(db *gorm.DB) (err error) {
	err = db.Delete(i).Error
	return
}
