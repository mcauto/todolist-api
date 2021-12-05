package todo

import (
	"todolist-api/modules/repository/_dbms"
)

// Service 인터페이스
type Service interface {
	Fetch(id uint64) (*Item, error)
	FetchAll(offset, limit int) ([]Item, error)
	Insert(item *Item) (int64, error)
	Update(id uint64, title string) (*Item, error)
	Delete(id uint64) (int64, error)
}

// ServiceImpl 구현체
type ServiceImpl struct {
	repo *_dbms.Repository
}

// NewService 생성자
func NewService(repo *_dbms.Repository) Service {
	return &ServiceImpl{repo}
}

// Fetch Fetch
func (s ServiceImpl) Fetch(id uint64) (item *Item, err error) {
	item = &Item{}
	tx := s.repo.First(item, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		item = nil
	}
	return item, tx.Error
}

// FetchAll FetchAll
func (s ServiceImpl) FetchAll(offset, limit int) (items []Item, err error) {
	items = make([]Item, limit)
	err = s.repo.Limit(limit).Offset(offset).Find(&items).Error
	return
}

// Insert Insert
func (s *ServiceImpl) Insert(item *Item) (int64, error) {
	tx := s.repo.Create(item)
	return tx.RowsAffected, tx.Error
}

// Update Update
func (s ServiceImpl) Update(id uint64, title string) (*Item, error) {
	item := &Item{}
	tx := s.repo.Model(item).Where("id = ?", id).Update("title", title)
	if tx.Error != nil || tx.RowsAffected == 0 {
		item = nil
	}
	return item, tx.Error
}

// Delete Delete
func (s ServiceImpl) Delete(id uint64) (int64, error) {
	tx := s.repo.Delete(&Item{ID: id})
	return tx.RowsAffected, tx.Error
}
