.PHONY: build
build: deps compile

.PHONY: deps
deps:
	@ go mod tidy --compat=1.20

.PHONY: compile
compile:
	@ go build -o bin/catgo cmd/main.go
