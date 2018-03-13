build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/vessel-service \
	proto/vessel/vessel.proto

	GOOS=linux GOARCH=amd64 go build 
		docker build -t vessel-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service