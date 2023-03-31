.PHONY: build
build:
	@buf generate && go build -o ./bin/protoc-gen-go-private ./plugin/...

.PHONY: clean
clean:
	@rm -rf ./bin


.PHONY: test
test: build
	@./scripts/test.sh
