.PHONY: build push

IMAGE = kubernetes-elasticsearch
TAG = 2.4.1-13

build:	elasticsearch_logging_discovery
	docker build -t mwrona/$(IMAGE):$(TAG) .

push:
	docker push mwrona/$(IMAGE):$(TAG)

elasticsearch_logging_discovery:
	go build -a -ldflags "-w" elasticsearch_logging_discovery.go
