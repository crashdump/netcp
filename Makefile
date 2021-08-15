VERSION:=$(shell head -n1 VERSION)
GIT_HASH_SHORT:=$(shell git rev-parse --short HEAD)
UNAME := $(shell uname -s)
GOOGLE_APPLICATION_CREDENTIALS ?= "~/.gcp/netcp-it-service-account.json"
GOOGLE_CLOUD_PROJECT ?= cloudcopy-it

ifeq ($(UNAME), Darwin)
	# TODO: Ensure the below automatically select the right identity so it's not tied to my machine.
	CMD_SIGN := codesign --verify --force -vv -s "Apple Development: ap@cdfr.net (SQ9LA8W8BB)" dist/app/netcp
endif

$(info ----------------------------)
$(info > Netcp Platform)
$(info ----------------------------)

.PHONY: all
all: build-srv
build: build-srv build-cli
test: test-srv # test-ui

.PHONY: build-ui
build-ui:
	cd ui/ && \
	 yarn install && \
	 yarn build

.PHONY: build-srv
build-srv:
	go build -o ./dist/app/netcp-srv -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/srv/v1/*

.PHONY: build-cli
build-cli:
	go build -o ./dist/app/netcp -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/cli/v1/*
	$(CMD_SIGN)

.PHONY: build-swagger
build-swagger:
	swag init -o ./api -g cmd/srv/v1/*

.PHONY: run
run:
	export GO_ENV=local
	air & \
	cd ui/; yarn serve  & \
	wait

#.PHONY: test-ui
#test-ui:
#	cd ui/ && \
#	  yarn lint && \
#	  yarn test:unit

.PHONY: test-api
test-srv:
	go fmt ./...
	go vet ./...
	go test -v ./... -race -coverprofile=coverage.out -covermode=atomic

.PHONY: clean
clean:
	go mod tidy
	rm -f dist/ui/*
	touch dist/ui/.gitkeep
	rm -f dist/app/*
	touch dist/app/.gitkeep