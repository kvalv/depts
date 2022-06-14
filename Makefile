.PHONY: test build install serve-docs

build:
	go build

install:
	go install


generated.go: query.graphql
	go run github.com/Khan/genqlient

test:
	go test
	
serve-docs:
	docker run --rm -it -p 8000:8000 -v ${PWD}:/docs squidfunk/mkdocs-material

.PHONY: local-release
local-release:
	goreleaser check
	goreleaser release --rm-dist --snapshot
