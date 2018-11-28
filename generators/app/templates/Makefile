.PHONY: build fmt lint dev test vet godep install bench run fresh

SHELL := /bin/sh
PKG_NAME=$(shell basename `pwd`)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GINKGO := $(shell command -v ginkgo 2> /dev/null)

vendor-dependencies:
	govendor add +e

build: vet \
		test
		go build

doc:
	godoc -http=:6060

fmt:
	go fmt

lint:
	golint ./... | grep -v vendor

dev:
	DEBUG=* go get && govendor add +e && go run main.go

fresh:
	go get github.com/pilu/fresh && export ENVIRONMENT=local && fresh

# Runs Tests with single stream logs for ginkgo tests
test:
# Install ginkgo if not installed
ifeq ($(GINKGO),)
	go get -u github.com/onsi/ginkgo/ginkgo
endif
	# Run all Gin tests with ginkgo
	export ENVIRONMENT=testing && ginkgo -r -v -noisyPendings=false --nodes=1 -notify --progress --trace ./gin | grep -v vendor

# Runs all tests
test-all:
	export ENVIRONMENT=testing && go test -v ./... -cover | grep -v vendor

test-only:
	export ENVIRONMENT=testing && ginkgo -r -v -noisyPendings=false --nodes=1 -notify --progress --trace ./gin --focus="ONLY"

bench:
	go test ./... -bench=. | grep -v vendor

vet:
	go vet

commit:
	read -r -p "Commit message: " message; \
	git add .; \
	git commit -m "$$message" \

# note pushes to origin
push: build
	git push origin $(GIT_BRANCH)

cover:
	read -r -p "package to get coverage from: " package; \
	export ENVIRONMENT=testing && go test -v ./"$$package" --cover --coverprofile=coverage.out | grep -v vendor \

	go tool cover -html=coverage.out

# watches and runs ginkgo tests and notifies on failures
watch:

	export ENVIRONMENT=testing && ginkgo watch -r -v -noisyPendings=false --nodes=1 -notify --progress --trace ./gin | grep -v vendor

prune:

	git branch --merged | egrep -v "(^\*|staging)" | xargs git branch -d && git remote prune origin
