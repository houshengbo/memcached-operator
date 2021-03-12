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
