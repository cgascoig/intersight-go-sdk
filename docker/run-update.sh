#!/bin/bash

set -e

export PATH="${PATH}:$(go env GOPATH)/bin"

git config --global user.email "cgascoig@cisco.com"
git config --global user.name "Chris Gascoigne"

BRANCH="${BRANCH:-main}"

BASE_DIR="${PWD}"

git clone "https://pat:${GITHUB_PAT}@github.com/cgascoig/intersight-go-sdk.git"
cd intersight-go-sdk/scripts
git checkout "${BRANCH}"
pipenv sync
pipenv run ./update.py

VER=$(cat ${BASE_DIR}/intersight-go-sdk/OPENAPI_VERSION)

cd "${BASE_DIR}"
git clone "https://pat:${GITHUB_PAT}@github.com/cgascoig/intersight-go-example.git"
cd intersight-go-example
go get "github.com/cgascoig/intersight-go-sdk/intersight@v${VER}"
make
if [[ -n $(git status --porcelain) ]]
then 
    echo "repo is dirty"
    git commit -a -m "Update SDK to v${VER}"
    git push origin main
    (cd ../intersight-go-sdk/scripts/; pipenv run ./notify.py "Example project successfully updated to v${VER}")
else 
    (cd ../intersight-go-sdk/scripts/; pipenv run ./notify.py "Example project unchanged at v${VER}")
fi