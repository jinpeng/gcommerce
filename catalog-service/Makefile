build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/jinpeng/gcommerce/catalog-service \
	  proto/catalog/catalog.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t catalog-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns catalog-service

clean:
	rm catalog-service proto/catalog/catalog.pb.go
