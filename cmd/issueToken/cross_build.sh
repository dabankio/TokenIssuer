#!/usr/bin/env bash

VERSION=`git describe --tags --dirty`
BUILD=`date +%FT%T%z`

xgo \
--targets=windows/386,windows/amd64,darwin/386,darwin/amd64,linux/386,linux/amd64 \
-ldflags "-s -w -X main.Version=$VERSION -X main.BuildDate=$BUILD" \
./