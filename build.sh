#!/bin/sh

init()
{
  export PATH="$PATH:$(go env GOPATH)/bin"
  go install github.com/GeertJohan/go.rice/rice@latest
  cd cmds/authentication && rice embed-go
  cd ../..
}

init

build_for_platform()
{
  GOOS=$1 GOARCH=$2 CGO_ENABLED=0 go build -o builds/mcdonalds-$1-$2
}

mkdir -p builds

build_for_platform linux arm64
build_for_platform linux amd64
build_for_platform darwin arm64
build_for_platform darwin amd64
build_for_platform windows amd64

mv builds/mcdonalds-windows-amd64 builds/mcdonalds-windows-amd64.exe
