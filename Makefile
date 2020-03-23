GOPATH     := $(shell pwd)/src/go
PORT       := 8080
REDIS_HOST ?= localhost
REDIS_PORT ?= 6379

.PHONY: build-server-src
build-server-src: $(GOPATH)/bin $(GOPATH)/pkg
	mkdir -p $(GOPATH)/src/github.com/montenegrodr/brcovid19api/
	GOPATH=$(GOPATH) swagger generate server -f ./swagger.yaml -A brcovid19api -t $(GOPATH)/src/github.com/montenegrodr/brcovid19api/

$(GOPATH)/bin:
	mkdir -p $(GOPATH)/bin

$(GOPATH)/pkg:
	mkdir -p $(GOPATH)/pkg

.PHONY: dep-all
dep-all: $(GOPATH)/bin $(GOPATH)/pkg
	cd $($GOPATH) && \
	go get -u github.com/golang/dep/cmd/dep && \
	cd $(GOPATH)/src/github.com/montenegrodr/brcovid19api && \
	$(GOPATH)/bin/dep ensure

.PHONY: build
build:
	go build $(GOPATH)/src/github.com/montenegrodr/brcovid19api/cmd/brcovid19api-server

.PHONY: start
start:
	REDIS_HOST=$(REDIS_HOST) REDIS_PORT=$(REDIS_PORT) ./brcovid19api-server --port $(PORT)
