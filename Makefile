# Basic Go makefile

APPNAME=GoCraft
GOCMD=go
GOBUILD=$(GOCMD) build -o $(APPNAME)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTIDY=$(GOCMD) mod tidy


all: build


tidy:
	$(GOTIDY)
build: tidy
	$(GOBUILD) -v -tags release
test: tidy
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
dev: tidy
	$(GOBUILD) -v -gcflags=all="-N -l" -tags "debug hints"
	./$(APPNAME)