SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
# .DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

GO_PATH ?= $(shell go env GOPATH)

OPENAPI_JAR := ~/openapi-generator-cli.jar
GIT_REPO_ID := intersight-go-sdk
GIT_USER_ID := cgascoig
# Default OPENAPI_VERSION - ideally this should be set as an env var
OPENAPI_VERSION ?= 1.0.0

all: intersight/README.md
.PHONY: all

clean:
> rm -Rf intersight
.PHONY: clean

intersight/README.md: spec/intersight-openapi-v3.json openapi-generator-config.json Makefile
> mkdir -p $(@D)
> java -jar $(OPENAPI_JAR) generate -g go -c openapi-generator-config.json -i spec/intersight-openapi-v3.json -o intersight/ --git-repo-id $(GIT_REPO_ID) --git-user-id $(GIT_USER_ID) --additional-properties=packageVersion=$(OPENAPI_VERSION)
> $(GO_PATH)/bin/goimports -l -w intersight