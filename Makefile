# Переменные
BINARY_NAME=myapp
DOCKER_IMAGE=myapp-image
GO_BUILD_ENV=CGO_ENABLED=0
GO_FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Цели
.PHONY: all build  clean docker-build docker-up docker-down build-linux-arm build-windows

all:  build


build:
	$(GO_BUILD_ENV) go build -o $(BINARY_NAME) ./cmd/main.go



# Очистка
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe
	go clean


docker-build:
	docker compose build


docker-up:
	docker compose up -d


docker-down:
	docker compose down


build-linux-arm:
	$(GO_BUILD_ENV) GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-linux-arm ./cmd/main.go

# Сборка под Windows (amd64)
build-windows:
	$(GO_BUILD_ENV) GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe ./cmd/main.go