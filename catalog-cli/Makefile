build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t catalog-cli .

run:
	docker run -e MICRO_REGISTRY=mdns catalog-cli
