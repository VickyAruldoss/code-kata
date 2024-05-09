WORK_DIR := $(shell pwd)

run_local:
	air

pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...
	golangci-lint run -v ./... --timeout=180s

lint:
	golangci-lint run

cover_cli:
	go test ./...  -coverprofile=c.out
	go tool cover -html="c.out" 

pre-push:	cover_cli

clean:
	chmod -R +w ./.gopath vendor || true
