# todolist-api

## architecture

![clean architecture](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)

## commands

- config: 실행 전 설정 값을 확인용

## How to run

### go run

```bash
$ export $(grep -v '^#' .env | xargs)
$ go run main.go run
[Fx] PROVIDE    *config.Settings <= todolist-api/modules/config.NewSettings()
[Fx] PROVIDE    todo.Service <= todolist-api/modules/domains/todo.NewService()
[Fx] PROVIDE    *echo.Echo <= todolist-api/modules/delivery/web.NewServer()
[Fx] PROVIDE    *todo.Handler <= todolist-api/modules/delivery/web/v1/todo.NewHandler()
[Fx] PROVIDE    logger.Interface <= todolist-api/modules/repository/_dbms.NewLogger()
[Fx] PROVIDE    _dbms.MySQLDialector <= todolist-api/modules/repository/_dbms.NewMySQLDialector()
[Fx] PROVIDE    _dbms.SQLiteDialector <= todolist-api/modules/repository/_dbms.NewSQLiteDialector()
[Fx] PROVIDE    *_dbms.Repository <= todolist-api/modules/repository/_dbms.NewRepository()
[Fx] PROVIDE    fx.Lifecycle <= go.uber.org/fx.New.func1()
[Fx] PROVIDE    fx.Shutdowner <= go.uber.org/fx.(*App).shutdowner-fm()
[Fx] PROVIDE    fx.DotGraph <= go.uber.org/fx.(*App).dotGraph-fm()
[Fx] INVOKE             todolist-api/modules/delivery/web.registerHook()
[Fx] INVOKE             todolist-api/modules/delivery/web/v1/todo.BindRoutes()
[Fx] INVOKE             todolist-api/modules/delivery/cmd.glob..func2.1()

2021/12/04 19:45:19 /Users/deo/projects/go/src/todolist-api/vendor/gorm.io/driver/sqlite/migrator.go:32
[0.056ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="todos"

2021/12/04 19:45:19 /Users/deo/projects/go/src/todolist-api/modules/delivery/cmd/run.go:22
[0.044ms] [rows:-] SELECT * FROM `todos` LIMIT 1

2021/12/04 19:45:19 /Users/deo/projects/go/src/todolist-api/vendor/gorm.io/driver/sqlite/migrator.go:257
[0.026ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type = "index" AND tbl_name = "todos" AND name = "idx_todos_title"

2021/12/04 19:45:19 /Users/deo/projects/go/src/todolist-api/vendor/gorm.io/driver/sqlite/migrator.go:257
[0.023ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type = "index" AND tbl_name = "todos" AND name = "idx_todos_deleted"
[Fx] HOOK OnStart               todolist-api/modules/delivery/web.registerHook.func1() executing (caller: todolist-api/modules/delivery/web.registerHook)
[Fx] HOOK OnStart               todolist-api/modules/delivery/web.registerHook.func1() called by todolist-api/modules/delivery/web.registerHook ran successfully in 4.276µs
[Fx] RUNNING

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.6.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:5000
```
http://localhost:5000/swagger/index.html

### docker

```bash
$ make build-docker
$ docker run --rm -d ghcr.io/mcauto/todolist-api:latest
```
