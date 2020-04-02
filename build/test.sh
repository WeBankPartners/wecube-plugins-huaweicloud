#!/bin/sh
set -e -x
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
apk -U add gcc musl-dev

cd $(dirname $0)/..

go test -v 

