package repository_test

import (
	"database/sql/driver"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"todo-list-back/models"

	"todo-list-back/todo/repository"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func wrapRegex(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}
func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("mysql", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)

	return mock, gormDB.Set("gorm:update_column", true)
}
func checkMock(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func getRowsForTodoItems(items []*models.TodoItem) *sqlmock.Rows {
	var todoItemFieldNames = []string{"id",
		"text",
		"checked",
	}
	rows := sqlmock.NewRows(todoItemFieldNames)
	for _, item := range items {
		rows.AddRow(item.ID, item.Text, item.Checked)
	}
	return rows
}

func getRowWithFields(fields []driver.Value) *sqlmock.Rows {
	fieldNames := []string{}
	for i := range fields {
		fieldNames = append(fieldNames, fmt.Sprintf("f%d", i))
	}

	return sqlmock.NewRows(fieldNames).AddRow(fields...)
}

type testQueryFunc func(t *testing.T, m sqlmock.Sqlmock, conn *gorm.DB)

func TestQueries(t *testing.T) {
	funcs := []testQueryFunc{
		testFindAll,
		testCreate,
		testFindByID,
		testUpdate,
		testDelete,
	}
	for _, f := range funcs {
		f := f // save range var
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		funcName = filepath.Ext(funcName)
		funcName = strings.TrimPrefix(funcName, ".")

		t.Run(funcName, func(t *testing.T) {
			t.Parallel()
			m, db := newDB()

			defer checkMock(t, m)
			f(t, m, db)
		})
	}
}

func testFindAll(t *testing.T, m sqlmock.Sqlmock, db *gorm.DB) {
	var mockTodoItems []*models.TodoItem
	err := faker.FakeData(&mockTodoItems)
	assert.NoError(t, err)

	mockRows := getRowsForTodoItems(mockTodoItems)

	m.ExpectQuery(wrapRegex("SELECT * FROM `todo_items`")).
		WillReturnRows(mockRows)

	repo := repository.NewGormRepository(db)
	items, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, mockTodoItems, items)
}

func testCreate(t *testing.T, m sqlmock.Sqlmock, db *gorm.DB) {
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)
	mockTodoItem.Checked = true // gorm에서는 default 전달 시 query에 생성하지 않음
	query := "INSERT INTO `todo_items` (`id`,`text`,`checked`) " +
		"VALUES (?,?,?)"
	args := []driver.Value{
		mockTodoItem.ID,
		mockTodoItem.Text,
		mockTodoItem.Checked,
	}
	m.ExpectExec(wrapRegex(query)).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(0, 1))
	repo := repository.NewGormRepository(db)
	err = repo.Create(&mockTodoItem)
	assert.NoError(t, err)

}

func testFindByID(t *testing.T, m sqlmock.Sqlmock, db *gorm.DB) {
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)

	mockRow := getRowsForTodoItems([]*models.TodoItem{&mockTodoItem})
	query := "SELECT * FROM `todo_items` WHERE (id = ?)"
	m.ExpectQuery(wrapRegex(query)).
		WillReturnRows(mockRow)

	repo := repository.NewGormRepository(db)
	item, err := repo.FindByID(mockTodoItem.ID)
	assert.NoError(t, err)
	fmt.Println(item)

}

func testUpdate(t *testing.T, m sqlmock.Sqlmock, db *gorm.DB) {
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)

	mockTodoItem.Checked = true // gorm에서는 default 전달 시 query에 생성하지 않음
	// gorm의 query 생성 규칙은 알파벳 순서

	query := "UPDATE `todo_items` SET `checked` = ?, `text` = ?  WHERE `todo_items`.`id` = ?"
	m.ExpectExec(wrapRegex(query)).
		WithArgs(mockTodoItem.Checked, mockTodoItem.Text, mockTodoItem.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewGormRepository(db)
	err = repo.Update(&mockTodoItem)
	assert.NoError(t, err)

}
func testDelete(t *testing.T, m sqlmock.Sqlmock, db *gorm.DB) {
	var mockTodoItem models.TodoItem
	err := faker.FakeData(&mockTodoItem)
	assert.NoError(t, err)

	mockTodoItem.Checked = true // gorm에서는 default 전달 시 query에 생성하지 않음
	// gorm의 query 생성 규칙은 알파벳 순서

	query := "DELETE FROM `todo_items`  WHERE `todo_items`.`id` = ?"
	m.ExpectExec(wrapRegex(query)).
		WithArgs(mockTodoItem.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewGormRepository(db)
	err = repo.Delete(&mockTodoItem)
	assert.NoError(t, err)
}
