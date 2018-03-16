build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/consignment-service \
	  proto/consignment/consignment.proto

	CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go 
		docker build -t consignment-service .

run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=false \
		consignment-service
		