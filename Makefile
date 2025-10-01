GO       ?= go
REVIVE   ?= revive
BIN_NAME ?= field
VERSION  ?= $(shell git describe --tags)
PREFIX   ?= /usr/local/

BIN_FILE            = $(shell realpath -m "$(PREFIX)/bin/$(BIN_NAME)")
LICENSE_FILE        = $(shell realpath -m "$(PREFIX)/share/licenses/$(BIN_NAME)/LICENSE")
BASH_COMPLETION_DIR = $(shell realpath -m "$(PREFIX)/share/bash-completion/completions")
ZSH_COMPLETION_DIR  = $(shell realpath -m "$(PREFIX)/share/zsh/site-functions")
FISH_COMPLETION_DIR = $(shell realpath -m "$(PREFIX)/share/fish/vendor_completions.d")

-include Makefile.local

build:
	$(GO) build -trimpath -ldflags '-s -w -X main.Version=$(VERSION)' -o $(BIN_NAME)

install:
	install -Dm755 $(BIN_NAME) "$(BIN_FILE)"
	install -Dm644 LICENSE "$(LICENSE_FILE)"
	# Install completions
	mkdir -p "$(BASH_COMPLETION_DIR)" "$(ZSH_COMPLETION_DIR)" "$(FISH_COMPLETION_DIR)"
	$(BIN_NAME) _carapace bash > "$(BASH_COMPLETION_DIR)/$(BIN_NAME)"
	$(BIN_NAME) _carapace zsh  > "$(ZSH_COMPLETION_DIR)/_$(BIN_NAME)"
	$(BIN_NAME) _carapace fish > "$(FISH_COMPLETION_DIR)/$(BIN_NAME).fish"
	@echo "Completions installed for bash, zsh, and fish."

lint:
	$(GO) test -v ./...
	$(REVIVE) -config revive.toml -formatter friendly ./...

docs-dev:
	bun run dev
