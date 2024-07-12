REGISTRY?=ghcr.io/getporter/examples/images/whalesay
LATEST=$(REGISTRY):latest

build:
	docker build -t $(LATEST) .

publish: build
	docker push $(LATEST)
