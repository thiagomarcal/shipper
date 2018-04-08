build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/vessel-service \
	proto/vessel/vessel.proto

	CGO_ENABLED=0 GOOS=linux go build  -o bin/vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go  
		docker build -t vessel-service .

deploy:
	docker tag vessel-service:latest 419011600616.dkr.ecr.us-east-1.amazonaws.com/vessel-service:test
	docker push 419011600616.dkr.ecr.us-east-1.amazonaws.com/vessel-service:test
	kubectl set image deployment/vessel vessel-service=419011600616.dkr.ecr.us-east-1.amazonaws.com/vessel-service:test

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		vessel-service
