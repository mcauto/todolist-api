package todo_test

import (
	"context"
	"log"
	"os"
	"testing"
	"todolist-api/modules/_test"
	"todolist-api/modules/config"
	"todolist-api/modules/domains/todo"
	"todolist-api/modules/repository/_dbms"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func registerHook(lifecycle fx.Lifecycle, settings *config.Settings, repo *_dbms.Repository) {
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return os.Remove("todolist.db")
		},
	})
}

// TestModules test modules
var TestModules = fx.Options(
	config.Modules,
	_dbms.Modules,
	fx.Provide(todo.NewService),
	fx.Invoke(registerHook),
	fx.Invoke(func(repo *_dbms.Repository) {
		if err := repo.AutoMigrate(&todo.Item{}); err != nil {
			log.Fatal(err)
		}
	}),
)

func TestFetch(t *testing.T) {
	f := func(service todo.Service) {
		item, err := service.Fetch(1)
		assert.Error(t, err)
		assert.Nil(t, item)
		affected, err := service.Insert(&todo.Item{Title: "test"})
		assert.NoError(t, err)
		assert.NotZero(t, affected)
		item, err = service.Fetch(1)
		assert.NoError(t, err)
		assert.NotNil(t, item)
	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
	defer app.RequireStop()
}

func TestFetchAll(t *testing.T) {
	f := func(service todo.Service) {
		offset, limit := 0, 10
		items, err := service.FetchAll(offset, limit)
		assert.NoError(t, err)
		assert.Empty(t, items)

		affected, err := service.Insert(&todo.Item{Title: "test"})
		assert.NoError(t, err)
		assert.NotZero(t, affected)
		items, err = service.FetchAll(offset, limit)
		assert.NoError(t, err)
		assert.NotEmpty(t, items)
	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
	defer app.RequireStop()
}

func TestInsert(t *testing.T) {
	f := func(service todo.Service) {
		item := todo.Item{Title: "test"}
		affected, err := service.Insert(&item)
		assert.NoError(t, err)
		assert.NotZero(t, affected)

		duplicate := todo.Item{Title: "test"}
		affected, err = service.Insert(&duplicate)
		assert.Error(t, err)
		assert.Zero(t, affected)
	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
	defer app.RequireStop()
}

func TestUpdate(t *testing.T) {
	f := func(service todo.Service) {
		id, title := 1, "test"
		updated, err := service.Update(uint64(id), title)
		assert.NoError(t, err)
		assert.Nil(t, updated)
	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
	defer app.RequireStop()
}

func TestDelete(t *testing.T) {
	f := func(service todo.Service) {
		id := 1
		affected, err := service.Delete(uint64(id))
		assert.NoError(t, err)
		assert.Zero(t, affected)
	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
	defer app.RequireStop()
}
