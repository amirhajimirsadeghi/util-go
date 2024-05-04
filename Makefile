.PHONY: update

integration:
	go test -v ./...
	
test:
	go test -v -short ./...

lint:
	golangci-lint run ./...

update:
	go get -v -u ./...
	go mod tidy
