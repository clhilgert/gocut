BINARY_NAME=gocut
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin cmd/main/main.go
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux cmd/main/main.go
	GOARCH=amd64 GOOS=windows go build -o  ./target/${BINARY_NAME}-windows cmd/main/main.go

run: build
	./target/${BINARY_NAME}-darwin
