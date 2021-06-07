-include .env


VERSION := $(shell git describe --tags)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

## Init: Create Autovpn config files $HOME/autovpn
init:
	@echo "  >  Creating autovpn config file"
	sudo mkdir -p /etc/autovpn/confs

## build: Complie Golang files
build: 
	@echo "  >  Building binary..."
	cd $(PROJECTNAME) && go build $(LDFLAGS) -o $(PROJECTNAME)

## start: Start Autovpn
start: 
	@echo "  >  Starting autovpn..."
	cd $(PROJECTNAME) && sudo ./$(PROJECTNAME) start

## start: Start Autovpn
stats:
	@echo "  >  Print autovpn connections stats..."
	cd $(PROJECTNAME) && sudo ./$(PROJECTNAME) stats

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo



