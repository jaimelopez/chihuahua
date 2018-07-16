.PHONY: build docker-build docker-push

image 	:= "chihuahua"
version := "latest"

build:
	go build -ldflags '-extldflags "-static"' -o chihuahua

docker-build:
	docker build -t $(image):$(version) --no-cache .

docker-push:
	docker push $(image)