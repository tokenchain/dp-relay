#!/usr/bin/make -f
VERSION := v$(shell cat version)              # $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
GPG_SIGNING_KEY = ''
export GO111MODULE = on
define update_check
 sh update.sh
endef
# process build tags

# process linker flags


#BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
SHOWTIMECMD :=  date "+%Y/%m/%d H:%M:%S"

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
	gox -osarch="linux/amd64" -mod=readonly  -output build/linux/dprelay ./cmd/relay
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
preinstall: go.sum
	sudo go get github.com/mitchellh/gox
########################################
### Tools & dependencies
go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
.PHONY: all build install go.sum
