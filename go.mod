module github.com/rancher/rio

go 1.12

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.1
	github.com/docker/distribution => github.com/docker/distribution v2.7.1-0.20190104202606-0ac367fd6bee+incompatible
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	github.com/jetstack/cert-manager => github.com/rancher/cert-manager v0.7.0-rio.1
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190411180116-681f9ce8ac52
)

require (
	github.com/Azure/azure-sdk-for-go v31.1.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest v13.0.1+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.9.1 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20171113091838-e9091a26100e // indirect
	github.com/Masterminds/sprig v2.17.1+incompatible
	github.com/aokoli/goutils v1.1.0
	github.com/aws/aws-sdk-go v1.21.2 // indirect
	github.com/beorn7/perks v1.0.0 // indirect
	github.com/cenkalti/backoff v2.0.0+incompatible // indirect
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/containerd/console v0.0.0-20181022165439-0650fd9eeb50
	github.com/containerd/containerd v1.3.0-0.20190507210959-7c1e88399ec0
	github.com/davecgh/go-spew v1.1.1
	github.com/deislabs/smi-sdk-go v0.0.0-20190819154013-e53a9b2d8c1a
	github.com/docker/cli v0.0.0-20190723080722-8560f9e8cdad // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c
	github.com/docker/go-units v0.4.0
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/drone/envsubst v0.0.0-20171016184023-f4d1a8ef8670
	github.com/envoyproxy/go-control-plane v0.8.6 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.0.14 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/gdamore/tcell v0.0.0-20190319073105-ec71b09872d7
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/gogo/googleapis v1.2.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/gophercloud/gophercloud v0.2.0 // indirect
	github.com/gorilla/mux v1.7.2 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/jetstack/cert-manager v0.7.2
	github.com/knative/build v0.6.0
	github.com/knative/pkg v0.0.0-20190514205332-5e4512dcb2ca
	github.com/knative/serving v0.6.1
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/linkerd/linkerd2 v0.0.0-20190906005929-a4b30b49313a
	github.com/linkerd/linkerd2-proxy-init v1.2.0 // indirect
	github.com/markbates/inflect v1.0.4 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mattn/go-shellwords v1.0.5
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/moby/buildkit v0.6.0
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/onsi/ginkgo v1.8.0
	github.com/opencontainers/runc v1.0.1-0.20190307181833-2b18fe1d885e // indirect
	github.com/openshift/api v0.0.0-20190322043348-8741ff068a47 // indirect
	github.com/opentracing/opentracing-go v1.0.2 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/common v0.4.0 // indirect
	github.com/prometheus/procfs v0.0.0-20190425082905-87a4384529e0 // indirect
	github.com/rancher/axe v0.0.0-20190531011056-59fcf8b44147
	github.com/rancher/gitwatcher v0.4.1
	github.com/rancher/mapper v0.0.0-20190814232720-058a8b7feb99
	github.com/rancher/rdns-server v0.4.2
	github.com/rancher/wrangler v0.2.0
	github.com/rancher/wrangler-api v0.2.1-0.20190905172212-0efb8da386d3
	github.com/rivo/tview v0.0.0-20190319111340-8d5eba0c2f51
	github.com/rivo/uniseg v0.0.0-20190313204849-f699dde9c340 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	github.com/tektoncd/pipeline v0.4.0
	github.com/urfave/cli v1.20.1-0.20190203184040-693af58b4d51
	github.com/wercker/stern v0.0.0-20190705090245-4fa46dd6987f // indirect
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.23.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.2
	gotest.tools v2.2.0+incompatible
	istio.io/api v0.0.0-20190902114838-92e636208314
	istio.io/istio v0.0.0-20190412222632-d19179769183
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/cli-runtime v0.0.0-20190606211212-7b5a46666fe9
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/helm v2.13.1+incompatible // indirect
	k8s.io/klog v0.4.0
	k8s.io/kubernetes v1.14.3
	k8s.io/utils v0.0.0-20190801114015-581e00157fb1 // indirect
	sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
	sigs.k8s.io/yaml v1.1.0
)
