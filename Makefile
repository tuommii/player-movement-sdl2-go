CC=go

all: build run

build:
	$(CC) build -o bin/game src/*.go
run:
	./bin/game