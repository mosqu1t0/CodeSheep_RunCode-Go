.PHONY: build

FILENAME=CodeSheep_runcode-2.1
SERVER_IP=81.68.156.124

all: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "build/${FILENAME}"
	@echo "finish <3"

deploy: ./build/${FILENAME}
	scp ./build/${FILENAME} root@${SERVER_IP}:~/CodeSheep/


