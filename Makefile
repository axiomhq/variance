# TOOLCHAIN
GO      := CGO_ENABLED=0 go
CGO     := CGO_ENABLED=1 go
GOFLAGS := -buildvcs=false

# ENVIRONMENT
VERBOSE =

# MISC
COVERPROFILE := coverage.out

# TAGS
GO_TEST_TAGS := netgo

# FLAGS
GO_TEST_FLAGS = -race -coverprofile=$(COVERPROFILE) -tags='$(GO_TEST_TAGS)'

# DEPENDENCIES
GOMODDEPS = go.mod go.sum

# Enable verbose test output if explicitly set.
GOTESTSUM_FLAGS =
ifdef VERBOSE
	GOTESTSUM_FLAGS += --format=standard-verbose
endif

# Export environment variables for all subshells.
export

.PHONY: all
all: dep fmt lint test ## Run dep, fmt, lint and test

.PHONY: build
build: ## Build the library (no binary output)
	@echo ">> building library"
	@$(GO) build ./...

.PHONY: clean
clean: ## Remove build and test artifacts
	@echo ">> cleaning up artifacts"
	@rm -rf $(COVERPROFILE) dep.stamp

.PHONY: coverage
coverage: $(COVERPROFILE) ## Calculate the code coverage score
	@echo ">> calculating code coverage"
	@$(GO) tool cover -func=$(COVERPROFILE) | grep total | awk '{print $$3}'

.PHONY: dep-clean
dep-clean: ## Remove obsolete dependencies
	@echo ">> cleaning dependencies"
	@$(GO) mod tidy

.PHONY: dep-upgrade
dep-upgrade: ## Upgrade all direct dependencies to their latest version
	@echo ">> upgrading dependencies"
	@$(GO) get $(shell $(GO) list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	@$(GO) get $(shell $(GO) list -f '{{range .Settings}}{{.Path}}{{"\n"}}{{end}}' tool)
	@$(MAKE) dep

.PHONY: dep-update
dep-update: dep-upgrade ## Alias for dep-upgrade

.PHONY: dep
dep: dep-clean dep.stamp ## Install and verify dependencies and remove obsolete ones

dep.stamp: $(GOMODDEPS)
	@echo ">> installing dependencies"
	@$(GO) mod download
	@$(GO) mod verify
	@touch $@

.PHONY: fmt
fmt: ## Format and simplify the source code using golangci-lint fmt
	@echo ">> formatting code"
	@$(GO) tool golangci-lint fmt

.PHONY: generate
generate: ## Generate code
	@echo ">> generating code"
	@$(GO) generate ./...

.PHONY: lint
lint: ## Lint the source code
	@echo ">> linting code"
	@$(GO) tool golangci-lint run

.PHONY: test
test: ## Run all unit tests. Run with VERBOSE=1 to get verbose test output
	@echo ">> running tests"
	@$(CGO) tool gotestsum $(GOTESTSUM_FLAGS) -- $(GO_TEST_FLAGS) ./...

.PHONY: test-short
test-short: ## Run unit tests in short mode
	@echo ">> running short tests"
	@$(CGO) tool gotestsum $(GOTESTSUM_FLAGS) -- $(GO_TEST_FLAGS) -short ./...

.PHONY: test-bench
test-bench: ## Run benchmarks
	@echo ">> running benchmarks"
	@$(GO) test -benchmem -bench=. ./...

.PHONY: test-race
test-race: ## Run tests with race detector
	@echo ">> running tests with race detector"
	@$(CGO) test -race ./...

.PHONY: mod-tidy
mod-tidy: dep-clean ## Tidy module dependencies

.PHONY: verify
verify: dep fmt lint test ## Run all verification steps

.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

# MISC TARGETS

$(COVERPROFILE):
	@$(MAKE) test
