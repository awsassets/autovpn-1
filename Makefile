-include .env


VERSION := $(shell git describe --tags)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

init:
	@echo "  >  Creating autovpn config file"
	mkdir $(HOME)/.autovpn

build: init
	cd cmd/
	@echo "  >  Building binary..."
	go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo



