.DEFAULT_GOAL := run

fmt:
	go fmt test.go
.PHONY: fmt

run: fmt
	go run test.go
.PHONY: run
