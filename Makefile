# TOOLCHAIN
GO	  := CGO_ENABLED=0 go
CGO	  := CGO_ENABLED=1 go
GOFMT := $(GO)fmt

# ENVIRONMENT
VERBOSE =

# GO TOOLS
GOTOOLS := $(shell cat tools.go | grep "_ \"" | awk '{ print $$2 }' | tr -d '"')

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

# FUNCTIONS
# func go-run-tool(name)
go-run-tool = $(CGO) run -mod=mod $(shell echo $(GOTOOLS) | tr ' ' '\n' | grep -w $1)

.PHONY: all
all: ## Run dep, fmt, lint and test
all: dep fmt lint test

.PHONY: clean
clean: ## Remove build and test artifacts
clean:
	@echo ">> cleaning up artifacts"
	@rm -rf $(COVERPROFILE) dep.stamp

.PHONY: coverage
coverage: ## Calculate the code coverage score
coverage: $(COVERPROFILE)
	@echo ">> calculating code coverage"
	@$(GO) tool cover -func=$(COVERPROFILE) | grep total | awk '{print $$3}'

.PHONY: dep-clean
dep-clean: ## Remove obsolete dependencies
dep-clean:
	@echo ">> cleaning dependencies"
	@$(GO) mod tidy

.PHONY: dep-upgrade
dep-upgrade: ## Upgrade all direct dependencies to their latest version
dep-upgrade:
	@echo ">> upgrading dependencies"
	@$(GO) get -d $(shell $(GO) list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	@make dep

.PHONY: dep-upgrade-tools
dep-upgrade-tools: ## Upgrade all tool dependencies to their latest version
dep-upgrade-tools: $(GOTOOLS)

.PHONY: dep
dep: ## Install and verify dependencies and remove obsolete ones
dep: dep-clean dep.stamp

dep.stamp: $(GOMODDEPS)
	@echo ">> installing dependencies"
	@$(GO) mod download
	@$(GO) mod verify
	@touch $@

.PHONY: fmt
fmt: ## Format and simplify the source code using `gofmt`
fmt:
	@echo ">> formatting code"
	@! $(GOFMT) -s -w $(shell find . -path -prune -o -name '*.go' -print) | grep '^'

.PHONY: lint
lint: ## Lint the source code
lint:
	@echo ">> linting code"
	@$(call go-run-tool, golangci-lint) run

.PHONY: test
test: ## Run all unit tests. Run with VERBOSE=1 to get verbose test output ('-v' flag).
test:
	@echo ">> running tests"
	@$(call go-run-tool, gotestsum) $(GOTESTSUM_FLAGS) -- $(GO_TEST_FLAGS) ./...

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# MISC TARGETS

$(COVERPROFILE):
	@make test
