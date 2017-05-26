.PHONY:	build push

IMAGE = kubernetes-es-fluentd
TAG = 1.11

build:
	docker build -t mwrona/$(IMAGE):$(TAG) .

push:
	docker push mwrona/$(IMAGE):$(TAG)
