build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/thiagomarcal/shipper/consignment-service \
	  proto/consignment/consignment.proto

	CGO_ENABLED=0 GOOS=linux go build  -o bin/consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go 
		docker build -t consignment-service .

deploy:
	docker tag consignment-service:latest 419011600616.dkr.ecr.us-east-1.amazonaws.com/consignment-service:latest
	docker push 419011600616.dkr.ecr.us-east-1.amazonaws.com/consignment-service:latest
	kubectl delete deployments consignment
	kubectl create -f deployments/deployment.yml


run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e DISABLE_AUTH=false \
		consignment-service
		