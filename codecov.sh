#!/usr/bin/env bash

set -e
echo "" > coverage.txt

go test ./... -race -coverprofile=profile.out -covermode=atomic
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi