SRC = ./**/**.go
BINARY = bin/webapp

start:
	go run ${SRC}

build:
	@go test ./...
	@go build -o ${BINARY} ${SRC}

test:
	go test ./...

clean:
	rm ${BINARY}
