VERSION:=$(shell head -n1 VERSION)
GIT_HASH_SHORT:=$(shell git rev-parse --short HEAD)
UNAME := $(shell uname -s)
GOOGLE_APPLICATION_CREDENTIALS ?= "~/.gcp/netcp-it-service-account.json"
GOOGLE_CLOUD_PROJECT ?= "cloudcopy-it"
GCLOUD_PROJECT ?= "cloudcopy-it"
FIRESTORE_EMULATOR_HOST := "localhost:8080"
FIREBASE_AUTH_EMULATOR_HOST := "localhost:9099"
FIREBASE_STORAGE_EMULATOR_HOST := "localhost:9199"

ifeq ($(UNAME), Darwin)
	# TODO: Ensure the below automatically select the right identity so it's not tied to my machine.
	CMD_SIGN := codesign --verify --force -vv -s "Apple Development: ap@cdfr.net (SQ9LA8W8BB)" dist/app/netcp*
endif

.PHONY: all
all: build-srv
test: test-srv # test-ui

.PHONY: build-ui
build-ui:
	cd ui/ && \
	 yarn install && \
	 yarn build

.PHONY: build
build:
	$(MAKE) build-srv
	$(MAKE) build-cli
	$(CMD_SIGN)
	$(MAKE) build-ui

.PHONY: build-srv
build-srv:
	go build -o ./dist/app/netcp-srv -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/srv/v1/main.go

.PHONY: build-cli
build-cli:
	go build -o ./dist/app/netcp -ldflags "-X main.Version=${VERSION}-${GIT_HASH_SHORT}" ./cmd/cli/v1/main.go

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
	go test -v ./... -coverprofile=coverage.out -covermode=atomic

.PHONY: clean
clean:
	go mod tidy
	rm -f dist/ui/*
	touch dist/ui/.gitkeep
	rm -f dist/app/*
	touch dist/app/.gitkeep