ENV = $(shell go env GOPATH)
GO_VERSION = $(shell go version)
GO111MODULE=on

# Look for versions prior to 1.10 which have a different fmt output
# and don't lint with gofmt against them.
ifneq (,$(findstring go version go1.8, $(GO_VERSION)))
	FMT=
else ifneq (,$(findstring go version go1.9, $(GO_VERSION)))
	FMT=
else
    FMT=--enable gofmt
endif

lint: # @HELP lint files and format if possible
	@echo "executing linter"
	gofmt -s -w .
	GO111MODULE=on golangci-lint run -c .golangci-lint.yml $(FMT) ./...
dep-linter: # @HELP install the linter dependency
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(ENV)/bin $(GOLANG_CI_VERSION)
build: # @HELP build the packages
	sh $(PWD)/scripts/build.sh
run-dennis-gateway:
	echo "running the api server"
	chmod +x ./scripts/run-server-grpc.sh
	./scripts/run-server-grpc.sh
config-up:
	sudo docker-compose up -d
config-down:
	docker-compose down
generate-proto:
	protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb -I proto --go-grpc_opt=paths=source_relative proto/*.proto

ci: # @HELP executes on CI
ci: deps test fuzz dep-linter lint

gh-ci: # @HELP executes on GitHub Actions
gh-ci: deps test dep-linter lint

all: deps test fuzz lint