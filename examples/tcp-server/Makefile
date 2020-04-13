.PHONY: format vet clean

prepare: format vet

all: prepare docker

format:
	go fmt

vet:
	go vet

docker:
	docker build -t dev/server .

run:
	docker run --rm -p 3333:3333 dev/server

clean:
	docker rmi dev/server
