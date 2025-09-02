.PHONY: prepare build

prepare:
	go mod tidy

build:
	go build -o dist/codesearch main.go
