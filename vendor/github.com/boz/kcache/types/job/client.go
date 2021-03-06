package job

import (
	"github.com/boz/kcache/client"
	"k8s.io/client-go/kubernetes"
)

const resourceName = "jobs"

func NewClient(cs kubernetes.Interface, ns string) client.Client {
	scope := cs.BatchV1()
	return client.ForResource(scope.RESTClient(), resourceName, ns)
}
