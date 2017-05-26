.PHONY: build push

# The current value of the tag to be used for building and
# pushing an image to gcr.io
IMAGE = curator
TAG = 1.0.4

build:
	docker push mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG) .

push:
	docker push mwrona/kubernetes-elasticsearch/$(IMAGE):$(TAG)
