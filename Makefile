.PHONY: build run

build:
	go build -o conntracker cmd/conntracker/main.go

run: build
	./conntracker
