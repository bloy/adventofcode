GOCMD=go
GOBUILD=${GOCMD} build
GORUN=${GOCMD} run
GOCLEAN=${GOCMD} clean
GOTEST=${GOCMD} test
GOGET=${GOCMD} get

default: all

all: build

build:
	$(GOBUILD) ./...

run:
	$(GORUN) ./...

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)

.phony: all build clean test default run
