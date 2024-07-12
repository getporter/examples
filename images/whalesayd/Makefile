REGISTRY?=ghcr.io/getporter/examples/images/whalesayd
TAG?=v0.1.1
VERSION=$(REGISTRY):$(TAG)
LATEST=$(REGISTRY):latest

build:
	GOOS=linux CGO_ENABLED=0 go build
	docker build -t $(LATEST) .

publish: build
	docker tag $(LATEST) $(VERSION)
	docker push $(VERSION)
	docker push $(LATEST)