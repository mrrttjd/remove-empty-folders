.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

buildq:
	go build -ldflags "-s -w" .
.PHONY:buildq

build: vet
	gox -osarch "linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64" \
	-ldflags "-s -w" \
	-tags "remove-empty-folders"    \
	-gocmd go        \
	-output "pkg/{{.OS}}_{{.Arch}}/remove-empty-folders"  \
	.
.PHONY:build
