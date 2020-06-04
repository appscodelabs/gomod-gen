#!/bin/bash
set -xeou pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")
pushd $SCRIPT_ROOT
pwd

cp go.mod.orig go.mod
rm -rf go.sum
go mod tidy || true

pop
