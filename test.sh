#!/bin/bash

set -e

quines=$(find quine* -maxdepth 0 -type d)

for quine in $quines; do
  echo "testing: $quine"
  diff -u -a $quine/main.go <(go run $quine/main.go)
  echo "passed"
done
