package mocks

import (
	"todo-list-back/models"

	"github.com/stretchr/testify/mock"
)

// Usecase mock struct
type Usecase struct {
	mock.Mock
}

// FindAll mock function
func (_m *Usecase) FindAll() ([]*models.TodoItem, error) {
	ret := _m.Called()
	var r0 []*models.TodoItem
	if rf, ok := ret.Get(0).(func() []*models.TodoItem); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]*models.TodoItem)
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
func (_m *Usecase) Create(item *models.TodoItem) error {
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
func (_m *Usecase) FindByID(id uint) (*models.TodoItem, error) {
	ret := _m.Called(id)
	var r0 *models.TodoItem
	if rf, ok := ret.Get(0).(func(uint) *models.TodoItem); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(*models.TodoItem)
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
func (_m *Usecase) Update(item *models.TodoItem) error {
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
func (_m *Usecase) Delete(item *models.TodoItem) error {
	ret := _m.Called(item)
	var r0 error
	if rf, ok := ret.Get(0).(func(*models.TodoItem) error); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
