BINARY_NAME = imageflow
DOCKER_IMAGE = imageflow:latest

.PHONY: build
build:
	go build -o $(BINARY_NAME) cmd/app/main.go

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	docker run --rm -p 8080:8080 $(DOCKER_IMAGE)

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
