.PHONY: test build clean

test:
	go test ./...

build:
	go build -o bin/go-samples ./cmd

clean:
	rm -rf bin/

fmt:
	go fmt ./...

vet:
	go vet ./...

