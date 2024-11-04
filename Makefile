.PHONY: build
build: 
	go mod tidy
	go build -o ./cmd/bin/mars-rover ./cmd/main.go

.PHONY: clean
clean:	
	rm -rf ./cmd/bin

.PHONY: test
test:
	go test ./...
