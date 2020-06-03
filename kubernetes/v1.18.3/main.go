package main

import (
	_ "k8s.io/api"
	_ "k8s.io/apiextensions-apiserver"
	_ "k8s.io/apimachinery"
	_ "k8s.io/apiserver"
	_ "k8s.io/cli-runtime"
	_ "k8s.io/client-go"
	_ "k8s.io/cloud-provider"
	_ "k8s.io/component-base"
	_ "k8s.io/kube-aggregator"
	_ "k8s.io/kubernetes"
	_ "k8s.io/kubectl"
	_ "kmodules.xyz/client-go"
	_ "kmodules.xyz/monitoring-agent-api"
	_ "kmodules.xyz/custom-resources"
	_ "kmodules.xyz/objectstore-api"
	_ "kmodules.xyz/offshoot-api"
	_ "kmodules.xyz/openshift"
	_ "kmodules.xyz/webhook-runtime"
	_ "kmodules.xyz/crd-schema-fuzz"
	_ "kmodules.xyz/prober"
)

func main() {}