module github.com/appscodelabs/gomod-tools/kubernetes/v1.18.3

go 1.14

require (
	k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/apiserver v0.18.3
	k8s.io/cli-runtime v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/cloud-provider v0.18.3
	k8s.io/component-base v0.18.3
	k8s.io/kube-aggregator v0.18.3
	k8s.io/kubectl v0.18.3
	k8s.io/kubernetes v1.18.3
	kmodules.xyz/client-go v0.0.0-20200525195850-2fd180961371
	kmodules.xyz/crd-schema-fuzz v0.0.0-20200521005638-2433a187de95
	kmodules.xyz/custom-resources v0.0.0-20200604135349-9e9f5c4fdba9
	kmodules.xyz/monitoring-agent-api v0.0.0-20200525002655-2aa50cb10ce9
	kmodules.xyz/objectstore-api v0.0.0-20200521103120-92080446e04d
	kmodules.xyz/offshoot-api v0.0.0-20200521035628-e135bf07b226
	kmodules.xyz/openshift v0.0.0-20200522123204-ce4abf5433c8
	kmodules.xyz/prober v0.0.0-20200521101241-adf06150535c
	kmodules.xyz/webhook-runtime v0.0.0-20200522123600-ca70a7e28ed0
)

replace (
	bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d
	git.apache.org/thrift.git => github.com/apache/thrift v0.13.0
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v35.0.0+incompatible
	github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.5.0
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.2.0
	github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.1.0
	github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.2.0
	github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.2.0
	github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0
	github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.1.0
	github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.5.0
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.0.0
	go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => github.com/kmodules/api v0.18.4-0.20200524125823-c8bc107809b9
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423
	k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200521033432-49d3646051ad
)
