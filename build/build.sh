#!/bin/sh
set -e -x 
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
apk -U add bash git  gcc musl-dev
cd $(dirname $0)/..
source $(dirname $0)/version.sh

LINKFLAGS="-linkmode external -extldflags -static -s"
go build -ldflags "-X main.VERSION=$VERSION $LINKFLAGS" 
