GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=uptimerobot
VERSION?=0.0.0

.PHONY: cmd
cmd:
	mkdir -p bin
	#GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY_NAME) .
	GO111MODULE=on $(GOCMD) build -o bin/$(BINARY_NAME) cmd/uptimerobot/main.go