build:
	GOOS=linux GOARCH=amd64 go build -o bin/user-cli
	docker build -t user-cli .

run:
	docker run --net="host" -e MICRO_REGISTRY=mdns user-cli