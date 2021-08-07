.PHONY: build

build:
		go build -v ./cmd/denbot

.DEFAULT_GOAL := build