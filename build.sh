#/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_VERSION="v0.0.1-alpha"
LATEST_UPDATE=$(git rev-parse HEAD)

if [ "$GOOS" = "windows" ]; then
    go build -ldflags "-H windowsgui -X main.buildTime=$BUILD_TIME -X main.buildVersion=$BUILD_VERSION -X main.latestUpdate=$LATEST_UPDATE" -o azlang.exe ./.
else
    go build -ldflags "-X main.buildTime=$BUILD_TIME -X main.buildVersion=$BUILD_VERSION -X main.latestUpdate=$LATEST_UPDATE" -o azlang ./
fi