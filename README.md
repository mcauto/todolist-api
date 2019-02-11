# Todo-list-back

Todo-list-back is todo list REST API by golang



## How to run

```bash
$ make
$ cd build && ./todo-list-back # Build and Run
```


## REST API test with curl

```bash
# CREATE todo item
$ curl -X POST -H 'Content-Type: application/json; charset=utf-8' -d '{"text":"learn echo framework", "checked":false}' http://127.0.0.1:1323/api/v1/todos
# READ todo list
$ curl -X GET http://127.0.0.1:1323/api/v1/todos/
# READ todo item
$ curl -X GET http://127.0.0.1:1323/api/v1/todos/1
# UPDATE todo item
$ curl -X PATCH -H 'Content-Type: application/json; charset=utf-8' -d '{"text":"learn echo framework", "checked": true}' http://127.0.0.1:1323/api/v1/todos/1
# DELETE todo item
$ curl -X DELETE http://127.0.0.1:1323/api/v1/todos/1
```



## Feature

- [x] Code quality management (gofmt, golint, goreportcard)
- [x] todo CRUD
- [ ] Swagger docs