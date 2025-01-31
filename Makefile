EXAMPLES := $(shell go list ./... | grep "examples")
UI_PATH := ./web/ui/gobench-ui
GITHASH := `git rev-parse HEAD`
GITTAG := `git describe --tags --always`
LDFLAGS="-X github.com/gobench-io/gobench/master.gitCommit=$(GITHASH) -X github.com/gobench-io/gobench/master.gitTag=$(GITTAG)"

.PHONY: lint build examples ent pb

pb:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		pb/agent.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		pb/executor.proto

# run supported packages
lint-pkgs:
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/staticcheck
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell

build-pkgs:
	GO111MODULE=off go get -u github.com/facebook/ent/cmd/entc

lint:
	$(exit $(go fmt ./... | wc -l))
	go vet ./...
	find . -type f -name "*.go" | xargs misspell -error -locale US
	staticcheck $(go list ./... | grep -v ent/privacy)

build:
	GOARCH=linux/arm64/v8 go build -ldflags $(LDFLAGS) -o gobench ./

test:
	./scripts/cov.sh

examples:
	$(foreach var, $(EXAMPLES), go build -buildmode=plugin -o ./.bin/${var}.so $(var);)

# generate ent models
ent:
	entc generate ./ent/schema

build-web-ui:
	cd $(UI_PATH) && yarn install && yarn run build-noauth

run:
	go run -ldflags $(LDFLAGS) .
