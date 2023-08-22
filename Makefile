SHELL := /bin/bash
DIRS=$(shell ls)
GO=go
NFPM=nfpm
NAME=kvm-agent

# root
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/. && pwd -P))
endif

BUILD_DIR := $(ROOT_DIR)/build

# output
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(BUILD_DIR)/output/
$(shell mkdir -p $(OUTPUT_DIR))
endif

# commit
GIT_COMMIT:=$(shell git rev-parse HEAD)

# version
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match="v*" --dirty | sed 's/-/./g')
endif

# buildfile
BUILDFILE = "./main.go"

.PHONY: all
all: clean build build-rpm build-deb

.PHONY: build
build:
	@echo "build===========> Building binary"
	@echo "build===========> $(shell go version)"
	@echo "build===========> Building binary $(OUTPUT_DIR) *[Git Info]: $(VERSION)-$(GIT_COMMIT)"
	@export CGO_ENABLED=0 && GOOS=linux go build -o $(OUTPUT_DIR)/$(NAME) -ldflags '-s -w' $(BUILDFILE)
	@echo "build===========> Build binary success"
	@cp -r $(ROOT_DIR)/config.yaml $(OUTPUT_DIR)/config.yaml
	@echo "build===========> Copy config.yaml success"

.PHONY: clean
clean:
	@echo "clean===========> Clean binary"
	@rm -rf $(OUTPUT_DIR)
	@echo "clean===========> Clean binary success"

.PHONY: build-rpm
build-rpm:
	@echo "build-pkg===========> Building rpm package"
	@$(NFPM) pkg -f $(BUILD_DIR)/nfpm.yaml --packager rpm --target $(BUILD_DIR)/tmp/
	@echo "build-pkg===========> Building rpm package success in $(BUILD_DIR)/tmp/"

.PHONY: build-deb
build-deb:
	@echo "build-pkg===========> Building deb package"
	@$(NFPM) pkg -f $(BUILD_DIR)/nfpm.yaml --packager deb --target $(BUILD_DIR)/tmp/
	@echo "build-pkg===========> Building deb package success in $(BUILD_DIR)/tmp/"