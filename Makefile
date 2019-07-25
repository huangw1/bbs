BINARY_NAME=bbs

all:gotool
	go build -o ${BINARY_NAME} .

clean:
	go clean -i

gotool:
	gofmt -w .
	go vet .

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run gofmt and go vet"
	@echo "make ca - generate ca files"

