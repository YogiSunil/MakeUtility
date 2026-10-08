build:
	go build -o apiforge .

test:
	go test ./...

bench:
	go test -bench=. ./...

vet:
	go vet ./...

.PHONY: build test bench vet
