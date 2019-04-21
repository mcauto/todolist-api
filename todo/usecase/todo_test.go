package usecase_test

import (
	"testing"
	"todo-list-back/models"
	"todo-list-back/todo/mocks"
	"todo-list-back/todo/usecase"

	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.Repository)
	var mockTodoItems []*models.TodoItem
	err := faker.FakeData(&mockTodoItems)
	assert.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindAll").Return(mockTodoItems, nil)
		useCase := usecase.NewTodoUsecase(mockRepo)
		items, err := useCase.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, mockTodoItems, items)
	})
}
func TestCreate(t *testing.T) {
	mockRepo := new(mocks.Repository)
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.AnythingOfType("*models.TodoItem")).Return(nil)
		useCase := usecase.NewTodoUsecase(mockRepo)
		err := useCase.Create(&mockTodoItem)
		assert.NoError(t, err)
	})
}
func TestFindByID(t *testing.T) {
	mockRepo := new(mocks.Repository)
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("FindByID", mock.AnythingOfType("uint")).Return(&mockTodoItem, nil)
		useCase := usecase.NewTodoUsecase(mockRepo)
		item, err := useCase.FindByID(mockTodoItem.ID)
		assert.NoError(t, err)
		assert.Equal(t, &mockTodoItem, item)
	})
}
func TestUpdate(t *testing.T) {
	mockRepo := new(mocks.Repository)
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Update", mock.AnythingOfType("*models.TodoItem")).Return(nil)
		useCase := usecase.NewTodoUsecase(mockRepo)
		err := useCase.Update(&mockTodoItem)
		assert.NoError(t, err)
	})
}
func TestDelete(t *testing.T) {
	mockRepo := new(mocks.Repository)
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Delete", mock.AnythingOfType("*models.TodoItem")).Return(nil)
		useCase := usecase.NewTodoUsecase(mockRepo)
		err := useCase.Delete(&mockTodoItem)
		assert.NoError(t, err)

	})
}
