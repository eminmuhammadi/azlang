#/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_VERSION="v0.0.1-alpha"
GIT_REPO="https://github.com/eminmuhammadi/azlang"
LATEST_GIT_COMMIT=$(git rev-parse HEAD)
GO_VERSION=$(go version)

go build . -ldflags "-X  main.Version=$BUILD_VERSION -X main.BuildTime=$BUILD_TIME -X main.GitRepo=$GIT_REPO -X main.GitCommit=$LATEST_GIT_COMMIT"