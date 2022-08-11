run:
	go run './cmd/main.go'

build:
	go build './...'

test:
	go install github.com/jpoles1/gopherbadger@v2.4.0
	gopherbadger -md="readme.md" -png=false

secure:
	# or install it into ./bin/
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s
	./bin/gosec -exclude=G104 ./...

link:
	ln -s ${GOPATH}/src/github.com/onelogin/onelogin-go-sdk .