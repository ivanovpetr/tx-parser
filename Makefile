#! /usr/bin/make -f

# Project variables.
BUILD_FOLDER = ./dist

## build: Build the binary.
build:
	@echo Building Parser...
	@-mkdir -p $(BUILD_FOLDER) 2> /dev/null
	@go build -o $(BUILD_FOLDER)/tx-parser ./cmd/parser/main.go