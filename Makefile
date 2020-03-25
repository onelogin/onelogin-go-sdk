run:
	go run './cmd/main.go'

build:
	go build './...'

test:
	go test './...'

link:
	ln -s ${GOPATH}/src/github.com/onelogin/onelogin-go-sdk .