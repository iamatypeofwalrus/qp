#!/usr/bin/env sh
dist=dist/

mkdir $dist

GOOS=linux GOARCH=amd64 go build -o $dist/qp.linux.amd64
GOOS=darwin GOARCH=amd64 go build -o $dist/qp.darwin.amd64
