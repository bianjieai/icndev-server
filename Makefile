GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
# GOTEST=$(GOCMD) test
# GOGET=$(GOCMD) get
GO_MOD=$(GOCMD) mod
GO_ENV=$(GOCMD) env
BINARY_NAME=icndev-server
BINARY_UNIX=$(BINARY_NAME)-unix
export GO111MODULE = on
export GOSUMDB=off
export GIT_TERMINAL_PROMPT=1

all: get_vendor build

get_vendor:
	@rm -rf vendor/
	@echo "--> Downloading dependencies"
	$(GO_MOD) download
	$(GO_MOD) vendor

build:
	$(GOBUILD) -o $(BINARY_NAME) .