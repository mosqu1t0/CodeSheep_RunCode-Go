.PHONY: build

FILENAME=CodeSheep_runcode-1.0
SERVER_IP=81.68.156.124

all: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "build/${FILENAME}"
	@echo "finish <3"

deploy: ./build/${FILENAME}
	cp -r work ./build/
	if [ -e ./build/codesheep.tar.gz ]; then rm ./build/codesheep.tar.gz; fi
	tar -zcvf ./build/codesheep.tar.gz ./build/*
	scp ./build/codesheep.tar.gz root@${SERVER_IP}:~/CodeSheep/


