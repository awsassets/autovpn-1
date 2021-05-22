-include .env


VERSION := $(shell git describe --tags)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

## Init: Create Autovpn config files $HOME/autovpn
init:
	@echo "  >  Creating autovpn config file"
	mkdir -p $(HOME)/.autovpn/confs

## build: Complie Golang files
build: 
	@echo "  >  Building binary..."
	cd cmd && go build $(LDFLAGS) -o $(PROJECTNAME)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo



