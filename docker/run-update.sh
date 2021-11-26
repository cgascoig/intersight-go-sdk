#!/bin/sh

export PATH="${PATH}:$(go env GOPATH)/bin"

git config --global user.email "cgascoig@cisco.com"
git config --global user.name "Chris Gascoigne"

BRANCH="${BRANCH:-main}"

git clone "https://pat:${GITHUB_PAT}@github.com/cgascoig/intersight-go-sdk.git"
cd intersight-go-sdk/scripts
git checkout "${BRANCH}"
pipenv sync
pipenv run ./update.py