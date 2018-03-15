build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/vessel-service \
	proto/vessel/vessel.proto

	CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go  
		docker build -t vessel-service .

run:
	docker run -p 50052:50051 \
	-e DB_HOST="datastore:27017" \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns vessel-service \