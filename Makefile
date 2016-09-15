NAME := picam2s3
LDFLAGS := -ldflags="-s -w"

GLIDE := $(shell command -v glide 2>&1 > /dev/null)

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): deps
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: build-raspi
build-raspi:
	GOOS=linux GOARCH=arm GOARM=6 go build $(LDFLAGS) -o bin/$(NAME)-raspi

.PHONY: deps
deps:
	glide install

.PHONY: glide
glide:
ifndef GLIDE
	curl https://glide.sh/get | sh
endif
