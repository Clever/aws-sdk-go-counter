include golang.mk

SHELL := /bin/bash
export PATH := $(PWD)/bin:$(PATH)
PKG = github.com/Clever/aws-sdk-go-counter
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /gen-go)

$(eval $(call golang-version-check,1.12))

.PHONY: all test build run $(PKGS) install_deps

all: test build

test: $(PKGS)
$(PKGS): golang-test-all-strict-deps
	$(call golang-test-all-strict,$@)

build:
	$(call golang-build,$(PKG)/example,bin/example)

run: build
	bin/example

install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)
