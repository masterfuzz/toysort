.PHONY: build test test-big format

NAME := toysort

clean:
	rm -r bin/ || true

build:
	go build -o bin/$(NAME) cmd/main.go

test:
	go test -v ./...


format:
	go fmt ./...

gen:
	go run testing/bigfile/gen.go

test-big:
	echo question.txt | go run cmd/main.go | cmp answer.txt - && echo Test passed

test-small:
	echo testing/small_question.txt | go run cmd/main.go -n 2 | cmp testing/small_answer.txt - && echo Test passed
