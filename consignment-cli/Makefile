build:
	GOOS=linux GOARCH=amd64 go build -o bin/consignment-cli
	docker build -t consignment-cli .

run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns consignment-cli \
	consignment.json \
	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMWM1N2UzZTgtODQwZi00NGFmLWEzOWQtZWY0YjY4Y2M0NmViIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoiZXdhbi52YWxlbnRpbmU4OUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCR3SFguVHdTSk9tSzIuRmVzZEduYzZlS1BXWjFyQnJIcWsva2h2THBzNm5IZDRjemtWb0Y0ZSJ9LCJleHAiOjE1MjE3MzQ1NTIsImlzcyI6InNoaXBwZXIudXNlciJ9.9qP0NHvUZazNhuXVrbsByhLeMYL6634tWOmEcCRqgD0

	