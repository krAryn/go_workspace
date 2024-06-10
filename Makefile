.DEFAULT_GOAL := run

fmt:
	go fmt ./pattern_programming/test.go
.PHONY: fmt

run: fmt
	go run ./pattern_programming/test.go
.PHONY: run
