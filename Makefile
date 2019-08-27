build:
	go build -o multiply
docker-build:
	docker build -t multiply_http .
docker-run:
	docker run -p 3002:8001 multiply_http
