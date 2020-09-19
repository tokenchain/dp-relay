#!/usr/bin/make -f

VERSION := v$(shell cat version)
# $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
current_dir := $(shell pwd)
GPG_SIGNING_KEY = ''
COMPRESSED_NAME:="replay_centos_$(VERSION).tar.gz"
BUILD_TARGET := "~/build/linux"
export GO111MODULE = on

define update_check
 sh update.sh
endef

define compress_file
	cd ~/build/linux
	pwd # Prints ~/build/linux if cd succeeded
	ls -l -a
	tar -czf $(COMPRESSED_NAME) "dprelay"
	mv $(COMPRESSED_NAME) $(current_dir)/build/linux/$(COMPRESSED_NAME)
endef

# process build tags

# process linker flags
ldflags = \
    -X dprelay/x.Name="Darkpool Relay" \
	-X dprelay/x.Version=$(VERSION) \
	-X dprelay/x.Commit=$(COMMIT)

SHOWTIMECMD := date "+%Y/%m/%d H:%M:%S"
BUILD_FLAGS := -ldflags '$(ldflags)'

all: lint install
	OS=linux

build: go.sum update-git
	ifeq ($(OS),Windows_NT)
		go build -mod=readonly -o build/dprelay.exe ./cmd/relay
	else
		go build -mod=readonly -o build/dprelay ./cmd/relay
		go build -o build/dprelay ./cmd/relay
	endif

centos: update-git go.sum
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o build/linux/dprelay ./cmd/relay

install: go.sum
	go install -mod=readonly ./cmd/relay

sign-release:
	if test -n "$(GPG_SIGNING_KEY)"; then \
	  gpg --default-key $(GPG_SIGNING_KEY) -a \
	      -o SHA256SUMS.sign -b SHA256SUMS; \
	fi;
lint:
	go run ./cmd/relay

update-git: go.sum
	$(update_check)

fullbuild: install centos buildcompress

####====================
### Tools & dependencies
####====================
go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
preinstall: go.sum
	@sh init.sh

.PHONY: all build install go.sum




.ONESHELL: # Only applies to all target
buildcompress: 
	cd $(current_dir)/build/linux && tar -czf $(COMPRESSED_NAME) "dprelay"