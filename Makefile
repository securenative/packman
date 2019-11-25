# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install

PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

build: test
			$(GOBUILD) -o packman -v main.go
run:
			$(GORUN) .
test:
			$(GOTEST) -v ./...

install: build
			$(GOINSTALL) -v .

release: test
	$(foreach GOOS, $(PLATFORMS),\
    $(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o packman-$(GOOS)-$(GOARCH))))
