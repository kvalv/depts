.PHONY: test build install

build:
	go build

install:
	go install


generated.go: query.graphql
	go run github.com/Khan/genqlient

test:
	go test
	
