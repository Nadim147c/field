GO       ?= go
REVIVE   ?= revive
BIN_NAME ?= field
VERSION  ?= $(shell git describe --tags)
PREFIX   ?= /usr/local/

BIN_FILE        = $(shell realpath -m "$(PREFIX)/bin/$(BIN_NAME)")
LICENSE_DIR     = $(shell realpath -m "$(PREFIX)/share/licenses/$(BIN_NAME)")
LICENSE_FILE    = $(shell realpath -m "$(LICENSE_DIR)/LICENSE")
BASH_COMPLETION = $(shell realpath -m "$(PREFIX)/share/bash-completion/completions/$(BIN_NAME)")
FISH_COMPLETION = $(shell realpath -m "$(PREFIX)/share/fish/vendor_completions.d/$(BIN_NAME).fish")
ZSH_COMPLETION  = $(shell realpath -m "$(PREFIX)/share/zsh/site-functions/_$(BIN_NAME)")

-include Makefile.local

.PHONY: build install uninstall test

build:
	$(GO) build -trimpath -ldflags '-s -w -X main.Version=$(VERSION)' -o $(BIN_NAME)

install:
	install -Dm755 $(BIN_NAME) "$(BIN_FILE)"
	install -Dm644 LICENSE "$(LICENSE_FILE)"
	$(BIN_NAME) _carapace bash | install -Dm644 /dev/stdin "$(BASH_COMPLETION)"
	$(BIN_NAME) _carapace zsh  | install -Dm644 /dev/stdin "$(ZSH_COMPLETION)"
	$(BIN_NAME) _carapace fish | install -Dm644 /dev/stdin "$(FISH_COMPLETION)"

uninstall:
	@rm -vrf "$(BIN_FILE)" "$(LICENSE_DIR)" \
		"$(BASH_COMPLETION)" "$(ZSH_COMPLETION)" "$(FISH_COMPLETION)"

test:
	$(GO) test -v ./...
	$(REVIVE) -config revive.toml -formatter friendly ./...
