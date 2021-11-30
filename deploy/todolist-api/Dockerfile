FROM golang:1.17.3-alpine as build

RUN go install golang.org/x/lint/golint@latest \
    && go install github.com/securego/gosec/v2/cmd/gosec@latest \
    && go install github.com/axw/gocov/gocov@latest \
    && go install github.com/matm/gocov-html@latest \
    && go install github.com/AlekSi/gocov-xml@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest

RUN apk add make

WORKDIR /go/src/todolist-api

COPY ../.. .

RUN make build

FROM scratch

COPY --from=build /go/src/todolist-api/bin/todolist-api /usr/bin/todolist-api

CMD ["todolist-api", "run"]