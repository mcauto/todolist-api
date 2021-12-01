package todo

import (
	"todolist-api/modules/repository/_mysql"
)

// Service 인터페이스
type Service interface {
	Fetch(id uint64) (*Item, error)
	FetchAll(offset, limit int) ([]Item, error)
	Insert(item *Item) error
	Update(item *Item) error
	Delete(id uint64) error
}

// ServiceImpl 구현체
type ServiceImpl struct {
	repo *_mysql.Repository
}

// NewService 생성자
func NewService(repo *_mysql.Repository) Service {
	return &ServiceImpl{repo}
}

// Fetch Fetch
func (s ServiceImpl) Fetch(id uint64) (item *Item, err error) {
	item = &Item{}
	err = s.repo.First(item, id).Error
	return
}

// FetchAll FetchAll
func (s ServiceImpl) FetchAll(offset, limit int) (items []Item, err error) {
	items = make([]Item, limit)
	err = s.repo.Limit(limit).Offset(offset).Find(&items).Error
	return
}

// Insert Insert
func (s ServiceImpl) Insert(item *Item) error {
	return s.repo.Create(&item).Error
}

// Update Update
func (s ServiceImpl) Update(item *Item) error {
	return s.repo.Save(item).Error
}

// Delete Delete
func (s ServiceImpl) Delete(id uint64) error {
	return s.repo.Delete(&Item{ID: id}).Error
}
