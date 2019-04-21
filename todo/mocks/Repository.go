package mocks

import (
	"todo-list-back/models"

	"github.com/stretchr/testify/mock"
)

// Repository mock struct
type Repository struct {
	mock.Mock
}

// FindAll mock function
func (_m *Repository) FindAll() ([]*models.TodoItem, error) {
	ret := _m.Called()
	var r0 []*models.TodoItem
	if rf, ok := ret.Get(0).(func() []*models.TodoItem); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TodoItem)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Create mock function
func (_m *Repository) Create(item *models.TodoItem) error {
	ret := _m.Called(item)
	var r0 error
	if rf, ok := ret.Get(0).(func(*models.TodoItem) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// FindByID mock function
func (_m *Repository) FindByID(id uint) (*models.TodoItem, error) {
	ret := _m.Called(id)
	var r0 *models.TodoItem
	if rf, ok := ret.Get(0).(func(uint) *models.TodoItem); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TodoItem)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Update mock function
func (_m *Repository) Update(item *models.TodoItem) error {
	ret := _m.Called(item)
	var r0 error
	if rf, ok := ret.Get(0).(func(*models.TodoItem) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// Delete mock function
func (_m *Repository) Delete(item *models.TodoItem) error {
	ret := _m.Called(item)
	var r0 error
	if rf, ok := ret.Get(0).(func(*models.TodoItem) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
