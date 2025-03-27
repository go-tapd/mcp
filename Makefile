ALL_PROJECT_DIRS := $(shell find ./cmd -maxdepth 1 -mindepth 1 -type d | sort | uniq) # find all project directories

GO = go
BIN_DIR = $(CURDIR)/bin

.PHONY: init
init:
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

.PHONY: lint
lint:
	golangci-lint run
	@echo "✅ Linting completed"

.PHONY: fix
fix:
	golangci-lint run --fix
	@echo "✅ Lint fixing completed"

.PHONY: test
test:
	go test ./... -race
	@echo "✅ Testing completed"

.PHONY: fmt
fmt:
	gofmt -w -e "vendor" .
	@echo "✅ Formatting completed"

.PHONY: fumpt
fumpt:
	gofumpt -w -e "vendor" .
	@echo "✅ Formatting completed"

.PHONY: nilaway-install
nilaway-install:
	go install go.uber.org/nilaway/cmd/nilaway@latest

.PHONY: nilaway
nilaway:
	nilaway ./...
	@echo "✅ Nilaway completed"

.PHONY: build
build: $(ALL_PROJECT_DIRS:%=build/%)
build/%: DIR=$*
build/%:
	@echo 'building $(DIR)' \
		&& cd $(DIR) \
		&& $(GO) build -o $(BIN_DIR) ./...