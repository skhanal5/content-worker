APP_NAME := clip-farmer-workflow
SRC_DIR := ./cmd/$(APP_NAME)
BUILD_DIR := ./bin
GO_FILES := $(shell find . -type f -name '*.go')

# I am on windows
GOOS:=windows 
GOARCH:=amd64

.PHONY: all build run clean test fmt lint

all: build

# Note: the exe extension for windows
build: $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME).exe $(SRC_DIR)/main.go

# Note: the exe extension for windows
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME).exe

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