.PHONY: all clean build run

APP_NAME := main

all: clean build run

clean:
	rm -rf $(APP_NAME)

build:
	go build -o $(APP_NAME) .

run:
	./$(APP_NAME)