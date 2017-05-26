.PHONY: build push

# The current value of the tag to be used for building and
# pushing an image to gcr.io
IMAGE = kubernetes-es-curator
TAG = 1.0.4

build:
	docker build -t mwrona/$(IMAGE):$(TAG) .

push:
	docker push mwrona/$(IMAGE):$(TAG)
