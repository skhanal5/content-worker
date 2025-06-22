APP_NAME := clip-farmer-workflow
SRC_DIR := ./cmd/$(APP_NAME)
BUILD_DIR := ./bin
GO_FILES := $(shell find . -type f -name '*.go')

.PHONY: all build run clean test fmt lint server

all: build

build: clean $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/main.go

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Linting code..."
	@golangci-lint run

server:
	@echo "Starting server..."
	temporal server start-dev --ui-port 8080