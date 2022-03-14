BINARY_NAME=giex

build:
	go build -o bin/${BINARY_NAME} main.go

release:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows main.go

clean:
	go clean
	rm -f bin/${BINARY_NAME}
	rm -f bin/${BINARY_NAME}-linux
	rm -f bin/${BINARY_NAME}-linux-arm
	rm -f bin/${BINARY_NAME}-linux-arm64
	rm -f bin/${BINARY_NAME}-darwin
	rm -f bin/${BINARY_NAME}-windows
