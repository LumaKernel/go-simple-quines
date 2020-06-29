#!/bin/bash

set -e

quines=$(find . -maxdepth 1 -name "quine*" -type d)

for quine in $quines; do
  echo "testing: $quine"
  diff -u -a $quine/main.go <(go run $quine/main.go)
  echo "passed"
done
