GREEN=\n\033[1;32;40m
NC=\033[0m # No Color

PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -name '*.go')

# 라이브러리 설치
# tidy: 미사용 라이브러리 제거
# vendor: vendor에 라이브러리 설치
# graph: 설치된 라이브러리 그래프 확인
ref:
	@/bin/sh -c 'echo "${GREEN}[library를 vendor에 설치합니다.]${NC}"'
	@go mod tidy
	@go mod vendor -v
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
	@gosec -fmt=json -out=.public/sast/results.json ./...
	@gosec -fmt=html -out=.public/sast/index.html ./...
.PHONY: sast

# 테스트 시작
test: ref
	@/bin/sh -c 'echo "${GREEN}[테스트를 시작합니다.]${NC}"'
	@unset LANG LC_ALL LC_MESSAGES && go test -short ${PKG_LIST}
.PHONY: test

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
	@gocov test ./... | gocov-html > .public/coverage/index.html
	@gocov test ./... | gocov-xml > coverage.xml
	@gocov test ./... | gocov report
.PHONY: coverage

# go report card
# https://github.com/gojp/goreportcard
report: ref
	@/bin/sh -c 'echo "${GREEN}[go report card를 생성합니다]${NC}"'
	@goreportcard-cli -d src
.PHONY: report

# 빌드
build: ref
	@/bin/sh -c 'echo "${GREEN}[빌드를 시작합니다]${NC}"'
	@mkdir -p bin
	@go build -o bin/
	@ls -al bin
.PHONY: build

install: build
	@/bin/sh -c 'echo "${GREEN}[설치를 시작합니다]${NC}"'
	@mv bin/* ${GOPATH}/bin/
.PHONY:

clean:
	@rm -rf bin vendor
.PHONY: clean

# swagger docs
# go install github.com/swaggo/swag/cmd/swag@latest
docs: ref
	@/bin/sh -c 'echo "${GREEN}[generate swagger docs]${NC}"'
	@swag init
.PHONY: docs