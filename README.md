# todolist-api

## architecture

![clean architecture](https://raw.githubusercontent.com/bxcodec/go-clean-arch/master/clean-arch.png)

## commands

- config: 실행 전 설정 값을 확인용

## 실행방법

### go run

```bash
$ export $(grep -v '^#' .env | xargs)
$ go rum main.go run
```

### docker

```bash
$ make build-docker
$ docker run --rm -d ghcr.io/mcauto/todolist-api:latest
```
