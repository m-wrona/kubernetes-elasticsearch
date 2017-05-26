.PHONY:	build push

IMAGE = fluentd-es
TAG = 1.11

build:
	docker build -t mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG) .

push:
	docker push mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG)
