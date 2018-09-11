#!/usr/bin/env bash

VERSION=`git describe --tags --dirty`
BUILD=`date +%FT%T%z`

go build -ldflags "-s -w -X github.com/dabankio/TokenIssuer/cmd/issueToken.Version=$VERSION -X github.com/dabankio/TokenIssuer/cmd/issueToken.BuildDate=$BUILD"