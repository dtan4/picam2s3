NAME := picam2s3
LDFLAGS := -ldflags="-s -w"

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME):
	go build $(LDFLAGS) -o bin/$(NAME)
