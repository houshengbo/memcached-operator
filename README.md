Create the operator with v1alpha1 CRD

Run the commands

```
mkdir memcached-operator
```

```
operator-sdk init --domain example.com --repo github.com/example/memcached-operator
```

Go through the instructions at 
```
https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/
```

Build and push the image:

```
export USER=<name>
make docker-build docker-push IMG=docker.io/$USER/memcached-operator:v0.0.1
```

Deploy the operator
```
make deploy IMG=docker.io/$USER/memcached-operator:v0.0.1
```

Create the CR:
```
kubectl apply -f config/samples/cache_v1alpha1_memcached.yaml
```

Create the operator with v1beta1 CRD
The v1beta1 CRD has the same content as v1alpha1 CRD.

```
operator-sdk create api --group cache --version v1beta1 --kind Memcached
```

Only create the resource, not the controller, because we will change the existing one.
The principle is that each controller will only reconcile on one version of the CR.

Add the `+kubebuilder:storageversion` marker to indicate the storage version. We add it into
Memcached at v1beta1.

Create the conversion webhook
```
operator-sdk create webhook --conversion --version v1beta1 --kind Memcached --group cache --defaulting --programmatic-validation
```
