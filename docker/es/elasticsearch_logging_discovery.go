package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
)

func flattenSubsets(subsets []api.EndpointSubset) []string {
	ips := []string{}
	for _, ss := range subsets {
		for _, addr := range ss.Addresses {
			ips = append(ips, fmt.Sprintf(`"%s"`, addr.IP))
		}
	}
	return ips
}

func main() {
	flag.Parse()
	glog.Info("Kubernetes Elasticsearch logging discovery")

	cc, err := restclient.InClusterConfig()
	if err != nil {
		glog.Fatalf("Failed to make client: %v", err)
	}
	client, err := clientset.NewForConfig(cc)

	if err != nil {
		glog.Fatalf("Failed to make client: %v", err)
	}
	namespace := metav1.NamespaceSystem
	envNamespace := os.Getenv("NAMESPACE")
	if envNamespace != "" {
		if _, err := client.Core().Namespaces().Get(envNamespace, meta_v1.GetOptions{}); err != nil {
			glog.Fatalf("%s namespace doesn't exist: %v", envNamespace, err)
		}
		namespace = envNamespace
	}

	var elasticsearch *api.Service
	// Look for endpoints associated with the Elasticsearch loggging service.
	// First wait for the service to become available.
	for t := time.Now(); time.Since(t) < 5*time.Minute; time.Sleep(10 * time.Second) {
		elasticsearch, err = client.Core().Services(namespace).Get("elasticsearch-logging", meta_v1.GetOptions{})
		if err == nil {
			break
		}
	}
	// If we did not find an elasticsearch logging service then log a warning
	// and return without adding any unicast hosts.
	if elasticsearch == nil {
		glog.Warningf("Failed to find the elasticsearch-logging service: %v", err)
		return
	}

	var endpoints *api.Endpoints
	addrs := []string{}
	// Wait for some endpoints.
	count := 0
	for t := time.Now(); time.Since(t) < 5*time.Minute; time.Sleep(10 * time.Second) {
		endpoints, err = client.Core().Endpoints(namespace).Get("elasticsearch-logging", meta_v1.GetOptions{})
		if err != nil {
			continue
		}
		addrs = flattenSubsets(endpoints.Subsets)
		glog.Infof("Found %s", addrs)
		if len(addrs) > 0 && len(addrs) == count {
			break
		}
		count = len(addrs)
	}
	// If there was an error finding endpoints then log a warning and quit.
	if err != nil {
		glog.Warningf("Error finding endpoints: %v", err)
		return
	}

	glog.Infof("Endpoints = %s", addrs)
	fmt.Printf("discovery.zen.ping.unicast.hosts: [%s]\n", strings.Join(addrs, ", "))
}