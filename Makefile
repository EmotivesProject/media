.PHONY: lint test integration

lint:
	golangci-lint run

test:
	go test -v ./...

integration:
	go test -v -tags=integration ./...
