/*
Copyright AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	_ "k8s.io/api"
	_ "k8s.io/apiextensions-apiserver"
	_ "k8s.io/apimachinery/pkg/util/errors"
	_ "k8s.io/apiserver/pkg/util/feature"
	_ "k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/cloud-provider"
	_ "k8s.io/component-base/version"
	_ "k8s.io/kube-aggregator"
	_ "k8s.io/kubectl/pkg/util"
	_ "k8s.io/kubernetes/pkg/util/env"
	_ "kmodules.xyz/client-go/core/v1"
	_ "kmodules.xyz/crd-schema-fuzz"
	_ "kmodules.xyz/custom-resources/apis/appcatalog"
	_ "kmodules.xyz/monitoring-agent-api/api/v1"
	_ "kmodules.xyz/objectstore-api/api/v1"
	_ "kmodules.xyz/offshoot-api/api/v1"
	_ "kmodules.xyz/openshift/apis/apps/v1"
	_ "kmodules.xyz/prober/api/v1"
	_ "kmodules.xyz/webhook-runtime/apis/workload/v1"
)

func main() {}
