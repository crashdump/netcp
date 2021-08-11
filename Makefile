VERSION:=$(shell head -n1 VERSION)
GIT_HASH_SHORT:=$(shell git rev-parse --short HEAD)
UNAME := $(shell uname -s)

ifeq ($(UNAME), Darwin)
	# TODO: Ensure the below automatically select the right identity so it's not tied to my machine.
	CMD_SIGN := codesign --verify --force -vv -s "Apple Development: ap@cdfr.net (SQ9LA8W8BB)" dist/app/netcpapp
endif


$(info ----------------------------)
$(info > Netcp Platform)
$(info ----------------------------)

.PHONY: all
all: build-api
build: build-api build-cli
test: test-ui test-api

.PHONY: build-ui
build-ui:
	cd ui/ && \
	 yarn install && \
	 yarn build

.PHONY: build-api
build-api:
	go build -o ./dist/app/netcp-srv -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/api/v1/main.go

.PHONY: build-cli
build-api:
	go build -o ./dist/app/netcp -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/cli/v1/main.go
	$(CMD_SIGN)

.PHONY: build-swagger
build-swagger: build-api
	swagger generate spec -o ./docs/swagger/swagger.json

.PHONY: run
run:
	export GO_ENV=local
	air & \
	cd ui/; yarn serve  & \
	wait

.PHONY: test-ui
test-ui:
	cd ui/ && \
	  yarn test:unit

.PHONY: test-api
test-api:
	go fmt ./...
	go test -v ./... -race -coverprofile=coverage.out -covermode=atomic

.PHONY: lint-ui
lint-ui:
	cd ui/ && \
	  yarn lint

.PHONY: lint-api
lint-api:
	go fmt ./...
	golangci-lint run -v

.PHONY: clean
clean:
	go mod tidy
	rm -f dist/ui/*
	touch dist/ui/.gitkeep
	rm -f dist/app/*
	touch dist/app/.gitkeep