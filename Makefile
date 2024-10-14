.PHONY: build run

build:
	CGO_ENABLED=0 go build -o conntracker cmd/conntracker/main.go

run: build
	./conntracker
