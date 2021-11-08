REPO_NAME := "openapi-boilerplate"
REPO_PATH := "adenix"
REPO_HOST := "github.com"
PKG := "${REPO_HOST}/${REPO_PATH}/${REPO_NAME}"
PKG_LIST := $(shell go list ${PKG}/...)

.PHONY: generate test cover coverhtml clean build

generate:
	@find . -name '*.gen.go' -delete
	@go generate ./pkg/...

test:
	@go test -cover ./...

cover:
	@mkdir -p cover
	@echo 'mode: atomic' > cover/coverage.out
	@echo ${PKG_LIST} | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=cover/coverage.tmp {} && tail -n +2 cover/coverage.tmp >> cover/coverage.out' && rm cover/coverage.tmp
	@go tool cover -func=cover/coverage.out

coverhtml: cover
	@go tool cover -html=cover/coverage.out

clean:
	@rm -rf bin cover

build: clean
	@go build $(flags) -o bin/ ${PKG_LIST}
