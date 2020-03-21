GOPATH := $(shell pwd)

.PHONY: build-server-src
build-server-src:
	mkdir -p $(GOPATH)/src/github.com/montenegrodr/brcovid19api/
	GOPATH=$(GOPATH) swagger generate server -f ./swagger.yaml -A brcovid19api -t $(GOPATH)/src/github.com/montenegrodr/brcovid19api/

.PHONY: dep
dep:
	go get -u github.com/golang/dep/cmd/dep

.PHONY: build
build:
	go build $(GOPATH)/src/github.com/montenegrodr/brcovid19api/cmd/brcovid19api-server