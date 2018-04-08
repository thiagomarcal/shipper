build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/user-service \
	  proto/user/user.proto

	CGO_ENABLED=0 GOOS=linux go build  -o bin/user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go
		docker build -t user-service .

deploy:
	docker tag user-service:latest 419011600616.dkr.ecr.us-east-1.amazonaws.com/user-service:latest
	docker push 419011600616.dkr.ecr.us-east-1.amazonaws.com/user-service:latest
	kubectl set image deployment/user user-service=419011600616.dkr.ecr.us-east-1.amazonaws.com/user-service:latest

run:
	docker run -d --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service
		