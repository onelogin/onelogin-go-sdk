.PHONY: test lint check build clean docs install-tools

build:
	go build .

test:
	go test -v -race -coverprofile=coverage.out .

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install golang.org/x/tools/cmd/godoc@latest

lint: install-tools
	go vet .
	golangci-lint run .
	$(shell go env GOPATH)/bin/gosec -exclude=G104 .

check: lint test
	go mod tidy
	git diff --exit-code go.mod go.sum

docs: install-tools
	$(shell go env GOPATH)/bin/godoc -http=:6060

clean:
	go clean -cache
	rm -f coverage.out
