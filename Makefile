.DEFAULT_GOAL := help
BIN := plow

build: $(BIN) ## Build all

$(BIN): main.go
	go build -o $@ $^

.PHONY: test
test: ## Run tests
	go test -v .

.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
