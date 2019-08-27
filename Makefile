USERNAME=mahendrakariya
IMAGE=multiply_http
VERSION=0.0.2

build:
	go build -o multiply

docker-build: build
	docker build -t $(USERNAME)/$(IMAGE):$(VERSION) .

docker-run: docker-build
	docker run -p 3002:8001 --network host $(USERNAME)/$(IMAGE):$(VERSION)

docker-push: docker-build
	docker push $(USERNAME)/$(IMAGE):$(VERSION)
