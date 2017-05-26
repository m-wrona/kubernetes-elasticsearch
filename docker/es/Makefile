.PHONY: build push

IMAGE = elastic
TAG = 2.4.1-13

build:	elasticsearch_logging_discovery
	docker build -t mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG) .

push:
	docker push mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG)

elasticsearch_logging_discovery:
	go build -a -ldflags "-w" elasticsearch_logging_discovery.go
