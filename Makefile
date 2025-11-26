.PHONY: build test test-big format

NAME := toysort

clean:
	rm -r bin/ || true

build:
	go build -o bin/$(NAME) cmd/main.go

test:
	go test -v ./...

test-big:
	echo TODO

format:
	go fmt ./...
