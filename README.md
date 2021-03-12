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
make docker-build docker-push IMG=docker.io/$USER/memcached-operator:v0.0.2
```

Deploy the operator
```
make deploy IMG=docker.io/$USER/memcached-operator:v0.0.2
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

## Install the cert-manager for webhooks

```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.2.0/cert-manager.yaml
```

## Storage migration job

Run this job when you need to migrate the existing CRDs from the v1alpha1 version into v1beta1 version.

Install the tool [`ko`](https://github.com/google/ko) first.

Build the image for the job in the root directory:
```
ko resolve -f config/post-install -B -t 0.0.2
```

The image will be published at `docker.io/$USER/migrate:v0.0.2`.
Replace the value of the image field with this value in `config/post-install/storage-version-migrator.yaml`.

Make sure you have installed the operator with the CRDs at v1beta1 version. Locally run the following command to migrate
the storage version:
```
kubectl apply -f config/post-install
```
