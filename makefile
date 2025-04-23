APP_NAME := ssh-manager
VERSION ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "dev")

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(APP_NAME) main.go

docker-build:
	docker build --build-arg GIT_COMMIT=$(VERSION) -t $(APP_NAME):$(VERSION) .

docker-run:
	docker run --rm -it $(APP_NAME):$(VERSION)

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

.PHONY: build docker-build docker-run run clean