.PHONY: build

FILENAME=CodeSheep_runcode-1.0

all: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "build/${FILENAME}"

deploy: ./build/${FILENAME}
	cp -r work ./build/
	tar -zcvf codesheep.tar.gz ./build/*
	scp ./codesheep.tar.gz root@175.178.69.145:~/CodeSheep/

