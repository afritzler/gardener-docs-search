GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gardener-docs-search
BINARY_UNIX=$(BINARY_NAME)_linux_amd64
MAIN=cmd/main.go
IMAGE=afritzler/gardener-docs-search
TAG=latest

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) $(BINARY_NAME) -v $(MAIN)
		./$(BINARY_NAME)

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(MAIN)
image: 
		docker build -t $(IMAGE):$(TAG) .