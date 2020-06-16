# igor
K8s operator to orchestrate the creation of new site deployments.

## Prerequisites
* A local K8s cluster
* A local docker registry running on port 5000
* Helm
* Make
* Bash or bash-like shell

### CRD Applied to Your Cluster
In order to create WebSite resources, you will need to make K8s aware of this
new resource type.

Run the following command in a terminal from the root of this repo.

```
$ cat ./crd/website-crd.yaml | kubectl apply -f -
```

## Building
Run the following command in a terminal from the root of this repo. This
will build a docker image, tag the image, and push it to your local docker
registry.

```
$ make build-docker
```

## Running
Run the following command in a terminal from the root of this repo. This will
generate and apply the yaml files needed to create your cluster resources. Once
done, you should have a running instance of the igor deployment running in your
cluster.

```
$ helm template ./helm/igor/ | kubectl apply -f -
```
