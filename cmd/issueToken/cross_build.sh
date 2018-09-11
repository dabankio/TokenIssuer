#!/usr/bin/env bash

VERSION=`git describe --tags --dirty`
BUILD=`date +%FT%T%z`

xgo \
--targets=windows/386,windows/amd64,darwin/386,darwin/amd64,linux/386,linux/amd64 \
-ldflags "-s -w -X github.com/dabankio/TokenIssuer/cmd/issueToken.Version=$VERSION -X github.com/dabankio/TokenIssuer/cmd/issueToken.BuildDate=$BUILD" \
./