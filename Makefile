GOPATH          := $(shell pwd)/src/go
PORT            := 8080
REDIS_HOST      ?= localhost
REDIS_PORT      ?= 6379
SERVICE_VERSION := 1.0
FETCHER_VERSION := 1.0

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

.PHONY: build-docker-service
build-docker-service:
	docker build -t montenegrodr/brcovid19api-service:$(SERVICE_VERSION) -f Dockerfile-service .

.PHONY: run-fetcher
run-fetcher:
	REDIS_HOST=$(REDIS_HOST) REDIS_PORT=$(REDIS_PORT) python src/python/fetch-data.py

.PHONY: build-docker-fetcher
build-docker-fetcher:
	docker build -t montenegrodr/brcovid19api-fetcher:$(FETCHER_VERSION) -f Dockerfile-fetcher .
