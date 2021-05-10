all: brew install

brew:
	brew bundle --no-lock

install: brew
	git init
	pre-commit uninstall
	pre-commit install -f
	tfenv install
	go get -u golang.org/x/lint/golint

init:
	go mod init github.com/AlaskaAirlines/tfmodule_azure_actiongroup
	go mod tidy

test:
	go get -t ./...
	go test -v ./test/actionGroup_test.go	

.PHONY: brew install test init
