# TOOLCHAIN
GO	  := CGO_ENABLED=0 go
CGO	  := CGO_ENABLED=1 go

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

.PHONY: all
all: dep fmt lint test ## Run dep, fmt, lint and test

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
	@$(GO) get $(shell $(GO) list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all) $(shell $(GO) list tool)
	@$(MAKE) dep

.PHONY: dep
dep: dep-clean dep.stamp ## Install and verify dependencies and remove obsolete ones

dep.stamp: $(GOMODDEPS)
	@echo ">> installing dependencies"
	@$(GO) mod download
	@$(GO) mod verify
	@touch $@

.PHONY: fmt
fmt: ## Format and simplify the source code using `golangci-lint fmt`
	@echo ">> formatting code"
	@$(GO) tool golangci-lint fmt

.PHONY: lint
lint: ## Lint the source code
	@echo ">> linting code"
	@$(GO) tool golangci-lint run

.PHONY: test
test: ## Run all unit tests. Run with VERBOSE=1 to get verbose test output ('-v' flag).
	@echo ">> running tests"
	@$(CGO) tool gotestsum $(GOTESTSUM_FLAGS) -- $(GO_TEST_FLAGS) ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# MISC TARGETS

$(COVERPROFILE):
	@$(MAKE) test
