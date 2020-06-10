module github.com/litmuschaos/chaos-runner

go 1.13

require (
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jpillora/go-ogle-analytics v0.0.0-20161213085824-14b04e0594ef
	// github.com/litmuschaos/chaos-operator v0.0.0-20200502085045-ae0a262d3baa
	// github.com/litmuschaos/chaos-operator v0.0.0-20200610171242-8c0c626224c4
	github.com/litmuschaos/chaos-operator v0.0.0-20200610171242-8c0c626224c4
	github.com/litmuschaos/elves v0.0.0-20200502084504-13be0e6937ee
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/pkg/errors v0.9.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	// k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a

	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	sigs.k8s.io/controller-runtime v0.3.0 // indirect
)