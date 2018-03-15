build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/user-service \
	  proto/user/user.proto

	CGO_ENABLED=0 GOOS=linux go build  -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go
		docker build -t user-service .

run:
	docker run -p 50053:50051 \
		-e DB_HOST="database:27017" \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service \
		