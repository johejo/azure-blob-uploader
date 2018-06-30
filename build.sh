#! /usr/bin/env bash

readonly VERSION=1.0.0
readonly TARGET=azure-blob-uploader

# Build for *nix
for GOOS in linux darwin freebsd; do
    for GOARCH in amd64 386 arm; do
        go build
        strip ${TARGET}
        upx ${TARGET}
        tar cfv ${TARGET}-${GOOS}-${GOARCH}.tar.gz  --use-compress-prog=pigz ${TARGET}
    done
done

# Build for Windows
GOOS=windows
for GOARCH in amd64 386; do
    go build
    strip ${TARGET}
    upx ${TARGET}
    zip ${TARGET}-${GOOS}-${GOARCH}.zip ${TARGET}.exe
done

rm -f ${TARGET} ${TARGET}.exe
