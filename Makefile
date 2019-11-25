# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install

build: test
			$(GOBUILD) -o packman -v cmd/packman/main.go
run:
			$(GORUN) .
test:
			$(GOTEST) -v ./...

install: build
			$(GOINSTALL) -v .