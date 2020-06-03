#!/bin/bash
set -xeou pipefail

cp go.mod.orig go.mod
rm -rf go.sum
go mod tidy || true
go mod edit -json > go.mod.json
