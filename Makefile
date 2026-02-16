
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/subclub
MAIN_PATH=./cmd/subclub

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOCMD) run $(MAIN_PATH)/*.go server

lint:
	golangci-lint run

tidy:
	$(GOCMD) mod tidy

deps:
	$(GOCMD) mod download

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build       Build the project binary"
	@echo "  test        Run unit tests"
	@echo "  clean       Clean build artifacts"
	@echo "  run         Run the application server"
	@echo "  lint        Run golangci-lint"
	@echo "  tidy        Run go mod tidy"
	@echo "  deps        Download dependencies"
