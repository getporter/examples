# GKE Kubernetes + Exec Mixin Example

This is a sample Porter bundle that makes use of both the Kubernetes and Exec
mixins to apply manifests to a GKE Kubernetes cluster.
The Kubernetes mixin is used to apply Kubernetes manifests to an
existing Kubernetes cluster, creating an NGINX deployment and a service. The
Kubernetes mixin is also used to produce an output with the value of the
service's ClusterIP.  After the `kubernetes` mixin finishes, the `exec` mixin is
ued to echo the cluster IP of the service that was created. 

To use this bundle, you will need an existing Kubernetes cluster and a
kubeconfig file for use as a credential.

```
porter build
porter credentials generate
porter install -c gke-example
```
