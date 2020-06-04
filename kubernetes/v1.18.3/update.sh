#!/bin/bash

# Copyright AppsCode Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -xeou pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")

pushd $SCRIPT_ROOT
pwd

cp go.mod.orig go.mod
rm -rf go.sum
go mod tidy
go mod edit -require k8s.io/client-go@v0.18.3

popd
