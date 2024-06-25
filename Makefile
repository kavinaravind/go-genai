NAME ?= chatbot

.PHONY: build

## build: Build the binary
build: 
	go build -o $(NAME)