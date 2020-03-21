GOPATH := $(shell pwd)

.PHONY: build-server
build-server:
	GOPATH=$(GOPATH) swagger generate server -f ./swagger.yaml -A application-name -t src
