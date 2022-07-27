include .env
export

GREEN=\n\033[1;32;40m
NC=\033[0m # No Color
PKG_LIST := $(shell go list -f '{{.Dir}}'/... -m | grep -v .back | grep -v config)
GO_FILES := $(shell find . -name '*.go')
APP_NAME := $(shell head -1 go.mod | cut -d " " -f2)
VERSION := $(shell cat modules/config/settings.go| grep "AppVersion" | head -1 | cut -d " " -f3 | cut -d '"' -f2)
DIR = ${PWD}/deploy/${APP_NAME}
GIT_REMOTE_URL = ghcr.io/mcauto

# 라이브러리 설치
# tidy: 미사용 라이브러리 제거
# vendor: vendor에 라이브러리 설치
# graph: 설치된 라이브러리 그래프 확인
ref:
	@/bin/sh -c 'echo "${GREEN}[library를 vendor에 설치합니다.]${NC}"'
	@go mod tidy
	@go mod vendor
	@/bin/sh -c 'echo "${GREEN}[그래프 확인]${NC}"'
	@go mod graph
.PHONY: ref

# 코딩 스타일 분석
# https://github.com/golang/lint
# go install golang.org/x/lint/golint@latest
lint: ref
	@/bin/sh -c 'echo "${GREEN}[정적분석(golint)을 시작합니다.]${NC}"'
	@golint -set_exit_status ${PKG_LIST}
.PHONY: lint

# 정적 분석
vet: ref
	@/bin/sh -c 'echo "${GREEN}[정적분석(vet)을 시작합니다.]${NC}"'
	@go vet ./...
.PHONY: vet

# 보안 정적 분석 (SAST)
# go install github.com/securego/gosec/v2/cmd/gosec@latest
sast: ref
	@/bin/sh -c 'echo "${GREEN}[보안정적분석(gosec)을 시작합니다.]${NC}"'
	@mkdir -p .public/sast
	@gosec -fmt=html -out=.public/sast/index.html modules/...; gosec -fmt=json -out=.public/sast/results.json modules/...; gosec ./...
.PHONY: sast

# mock 코드 자동 생성
# https://github.com/vektra/mockery
# --dir mock 코드 생성이 필요한 인터페이스 탐색 시작 위치
# --inpackage 해당 코드 위치에 생성
# --testonly _test.go를 붙여서 coverage에서 제외 (다른 패키지에서 사용할 수 없다)
# --case underscore 파일 이름 포맷
# @unset LANG LC_ALL LC_MESSAGES && mockery --dir modules --inpackage --all --case underscore
# --packageprefix mock_
mocks: ref
	@/bin/sh -c 'echo "${GREEN}[테스트를 시작합니다.]${NC}"'
	@unset LANG LC_ALL LC_MESSAGES && mockery --dir modules --all --case underscore --keeptree
.PHONY: mock

# 테스트 시작
test: ref
	@/bin/sh -c 'echo "${GREEN}[테스트를 시작합니다.]${NC}"'
	@unset LANG LC_ALL LC_MESSAGES && go test -short ${PKG_LIST}
.PHONY: test

benchmark: ref
	@/bin/sh -c 'echo "${GREEN}[테스트를 시작합니다.]${NC}"'
	@unset LANG LC_ALL LC_MESSAGES && go test -race -benchmem -bench . ${PKG_LIST}
.PHONY: benchmark

# race condition 검사
race: ref
	@/bin/sh -c 'echo "${GREEN}[race condition을 검사합니다.]${NC}"'
	@unset LANG LC_ALL LC_MESSAGES && go test -race -v ${PKG_LIST}
.PHONY: race

# 테스트 커버리지
# go install github.com/axw/gocov/gocov@latest
# go install github.com/matm/gocov-html@latest
# go install github.com/AlekSi/gocov-xml@latest
coverage: ref
	@/bin/sh -c 'echo "${GREEN}[test coverage를 계산합니다.]${NC}"'
	@mkdir -p .public/coverage
	@gocov test ${PKG_LIST} | gocov-html > .public/coverage/index.html
	@gocov test ${PKG_LIST} | gocov-xml > coverage.xml
	@gocov test ${PKG_LIST} | gocov report
.PHONY: coverage

# go report card
# https://github.com/gojp/goreportcard
report: ref
	@/bin/sh -c 'echo "${GREEN}[go report card를 생성합니다]${NC}"'
	@goreportcard-cli -d modules
.PHONY: report

# 빌드
build: ref
	@/bin/sh -c 'echo "${GREEN}[빌드를 시작합니다]${NC}"'
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/
	@ls -al bin
.PHONY: build

# 문서 생성
# go install github.com/swaggo/swag/cmd/swag@latest
docs: ref
	@/bin/sh -c 'echo "${GREEN}[swagger 문서를 생성합니다]${NC}"'
	@swag init --exclude=modules/echo-apps/keyauth,modules/echo-apps/whitelist,modules/echo-apps/geoip/v1/version
.PHONY: docs

docs-dev: ref
	@/bin/sh -c 'echo "${GREEN}[swagger 문서를 생성합니다]${NC}"'
	@swag init
.PHONY: docs-dev

build-docker:
	@/bin/sh -c 'echo "${GREEN}[image 빌드를 시작합니다]${NC}"'
	@docker build \
		-f "${DIR}/Dockerfile" \
		-t ${GIT_REMOTE_URL}/${APP_NAME}:latest .
	@docker build \
		-f "${DIR}/Dockerfile" \
		-t ${GIT_REMOTE_URL}/${APP_NAME}:v${VERSION} .
.PHONY: build-docker

install: build
	@/bin/sh -c 'echo "${GREEN}[설치를 시작합니다]${NC}"'
	@mv bin/* ${GOPATH}/bin/
.PHONY:

run:
	@/bin/sh -c 'echo "${GREEN}[${APP_NAME} 실행]${NC}"'
	@docker compose -p ${APP_NAME} -f deploy/docker-compose.yml up --build -d
.PHONY: run

exit:
	@/bin/sh -c 'echo "${GREEN}[${APP_NAME} 종료]${NC}"'
	@docker compose -p ${APP_NAME} -f deploy/docker-compose.yml down
.PHONY: exit

clean:
	@rm -rf bin vendor
.PHONY: clean

current_changelog:
	@/bin/sh -c "echo \"${GREEN}[release version] $(shell npx standard-version --dry-run | grep tagging | cut -d ' ' -f4)${NC}\""
	@/bin/sh -c "echo \"${GREEN}[description] ${NC}\""
	@npx standard-version --dry-run --silent | grep -v Done | grep -v "\-\-\-" | grep -v standard-version
.PHONY: current_changelog
