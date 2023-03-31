.PHONY: dependencies
dependencies:
	./scripts/dependencies.sh

.PHONY: build
build:
	@buf generate && go build -o ./bin/protoc-gen-go-private ./plugin/...

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: test
test:  dependencies build
	@./scripts/test.sh

.PHONY: fmt
fmt:
	@go fmt ./...