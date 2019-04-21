package repository

import (
	"todo-list-back/models"
	"todo-list-back/todo"

	"github.com/jinzhu/gorm"
)

type gormRepository struct {
	Conn *gorm.DB
}

// NewGormRepository gorm을 이용한 저장소 구현
func NewGormRepository(conn *gorm.DB) todo.Repository {
	return &gormRepository{
		Conn: conn,
	}
}

// FindAll find todo item
func (repo *gormRepository) FindAll() (items []*models.TodoItem, err error) {
	err = repo.Conn.Find(&items).Error
	return
}

// Create create todo item
func (repo *gormRepository) Create(item *models.TodoItem) (err error) {
	err = repo.Conn.Create(item).Error
	return
}

// Find find todo item
func (repo *gormRepository) FindByID(id uint) (*models.TodoItem, error) {
	item := &models.TodoItem{}
	err := repo.Conn.Where("id = ?", id).Find(item).Error
	return item, err
}

// Update update todo item as checked
func (repo *gormRepository) Update(item *models.TodoItem) (err error) {
	u := map[string]interface{}{
		"Text":    item.Text,
		"Checked": item.Checked,
	}
	err = repo.Conn.Model(item).Update(u).Error
	return
}

// Delete delete todo item
func (repo *gormRepository) Delete(item *models.TodoItem) (err error) {
	err = repo.Conn.Delete(item).Error
	return
}
