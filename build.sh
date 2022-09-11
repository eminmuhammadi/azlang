#/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_VERSION="v0.0.1-alpha"
LATEST_UPDATE=$(git rev-parse HEAD)

go build -ldflags "-X  main.VERSION=$BUILD_VERSION -X main.BUILD_TIME=$BUILD_TIME -X main.LATEST_UPDATE=$LATEST_UPDATE" -o azlang ./.