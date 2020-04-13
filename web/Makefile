.PHONY: format vet clean

current_dir := $(notdir $(shell pwd))


all: format vet build

format:
	go fmt

vet:
	go vet

build:
	go build -o $(current_dir).o

clean:
	rm $(current_dir).o
