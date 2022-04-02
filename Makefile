GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags || echo "unknown version")
BUILDTIME=$(shell date -u)
GOBUILD=CGO_ENABLED=0 go build -trimpath

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	${GOBUILD} -o alfred-ssh-workflow *.go

.PHONY: run
run:
	export alfred_workflow_bundleid="com.ifooth.alfred-ssh-workflow" && \
	export alfred_workflow_cache="./.alfred/cache" && \
	export alfred_workflow_data="./.alfred/data" && \
    go run main.go

.PHONY: test
test:
	@go test -v ./... -cover
