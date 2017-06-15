default: help

help:
	@echo -e "Select a sub command \n"
	@echo -e "build: \n\t Build binary"
	@echo -e "fmt: \n\t Format codes"
	@echo -e "run: \n\t Run mesoshosts"
	@echo -e "help: \n\t Display this help"
	@echo -e "\n"
	@echo -e "See README.md for more."

build:
	go build -o bin/mesoshosts

fmt:
	go fmt ./...
run:
	./bin/mesoshosts
