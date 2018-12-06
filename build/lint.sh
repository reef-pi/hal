#!/bin/bash

set -e

cd "$(dirname $0)"
if [ -n "$(diff <(goimports -d .) <(printf ""))" ]; then
  echo "Imports are not sorted"
  exit 1
fi

if [ -n "$(gofmt -d -s ./)" ]; then
  echo "Code is not formatted with -s"
  exit 1
fi
