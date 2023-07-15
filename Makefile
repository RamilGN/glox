.PHONY: lint
lint:
	@golangci-lint run ./... --out-format tab

.PHONY: test
test:
	@go test ./...

.PHONY: gen
gen:
	@go generate ./...
