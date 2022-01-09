#!/bin/bash

set -e

DIRNAME="$(dirname ${BASH_SOURCE[0]})"
ROOT="$(cd $DIRNAME; pwd)"

main() {
  build_app
  build_bin
}

build_app() {
  pushd "${ROOT}/app"
  [ ! -d "node_modules" ] && npm install
  npm run build
  popd
}

generate() {
   pushd "${ROOT}"
   GOPATH= go generate ./...
   popd
}

build_bin() {
  pushd "${ROOT}"
  GOPATH= go build ./
  popd
}

main
