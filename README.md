# kubernetes-elasticsearch

Elasticsearch & Kibana configuration for Kubernetes.

***Note: All images have been customized based on Kubernetes source code.***

Repository contains:

* Ingress with `basic auth`. SSL support can be enabled by providing cert.

* Kibana

* ES & curator (for cleaning old logs)

* Fluentd (for pushing logs from Kubernetes to ES)

# Dockers

Images:

* Fluentd: gathers logs from each node and sends it to ES

* Elasticsearch: indexing

* Kibana: display

* Curator: plugin for ES to clean up automatically old logs

#### Build image

```
make build
```

#### Manually push image to registry

```
make push
```

***Note: image will be built automatically and push to docker registry after each commit on master or tag created***

# DevOps

TODO: create uber script for deployment

# Local development

Prerequisite:

* [Mini Kube](https://github.com/kubernetes/minikube)

In order to access Kibana using `minikube` execute:

```
minikube service -n elk ingress-es
```

# Access Elasticsearch & Kibana

1) Connect to Kubernetes cluster

```
kube-ctl proxy
```

2) Check address of Elasticsearch & Kibana

```
kube-ctl cluster-info
```

3) Basic auth credentials `admin:admin`
