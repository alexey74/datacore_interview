.PHONY: deps test build publish

all: deps test build

deps:
	mkdir -p ./output

test:
	go test -count=1 -coverprofile=./output/golang-coverage.out ./... | tee ./output/golang-test.out

build:
	go build -o ./output/golang-artefact ./cmd/takehome

publish:
	if [[ "$(git branch --show-current)" -ne "master" ]]; then exit 1; fi
