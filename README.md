# README

This is a Keptn Service Template written in GoLang. 

Quick start:

* Download this repo as a zip-file, extract it into a new folder named after the service you want to create (e.g., simple-service) 
* Replace every occurrence of (docker) image names and tags from `christiankreuzbergerdtx/wget-test-service` to your docker organization and image name (e.g., `yourorganization/simple-service`)
* Replace every occurrence of `wget-test-service` with the name of your service (e.g., `simple-service`)
* Push your code a Git repo (e.g., GitHub) and adapt all links that contain `github.com` (e.g., to `github.com/christiankreuzbergerdtx/simple-service`)
* Àdapt the [go.mod](go.mod) file and change `example.com/` to the actual package name (e.g., `github.com/christiankreuzbergerdtx/simple-service`)
* Add yourself to the [CODEOWNERS](CODEOWNERS) file
* Last but not least: Remove this intro in the README file 

# wget-test-service
![GitHub release (latest by date)](https://img.shields.io/github/v/release/christiankreuzbergerdtx/wget-test-service)
[![Build Status](https://travis-ci.org/christiankreuzbergerdtx/wget-test-service.svg?branch=master)](https://travis-ci.org/christiankreuzbergerdtx/wget-test-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/christiankreuzbergerdtx/wget-test-service)](https://goreportcard.com/report/github.com/christiankreuzbergerdtx/wget-test-service)

This implements a wget-test-service for Keptn.

## Compatibility Matrix

| Keptn Version    | [Simple-Test-Service Docker Image](https://hub.docker.com/r/christiankreuzbergerdtx/wget-test-service/tags) |
|:----------------:|:----------------------------------------:|
|       0.6.1      | christiankreuzbergerdtx/wget-test-service:0.1.0 |

## Installation

The *wget-test-service* can be installed as a part of [Keptn's uniform](https://keptn.sh).

### Deploy in your Kubernetes cluster

To deploy the current version of the *wget-test-service* in your Keptn Kubernetes cluster, apply the [`deploy/service.yaml`](deploy/service.yaml) file:

```console
kubectl apply -f deploy/service.yaml
```

This should install the `wget-test-service` together with a Keptn `distributor` into the `keptn` namespace, which you can verify using

```console
kubectl -n keptn get deployment wget-test-service -o wide
kubectl -n keptn get pods -l run=wget-test-service
```

### Up- or Downgrading

Adapt and use the following command in case you want to up- or downgrade your installed version (specified by the `$VERSION` placeholder):

```console
kubectl -n keptn set image deployment/wget-test-service wget-test-service=christiankreuzbergerdtx/wget-test-service:$VERSION --record
```

### Uninstall

To delete a deployed *wget-test-service*, use the file `deploy/*.yaml` files from this repository and delete the Kubernetes resources:

```console
kubectl delete -f deploy/service.yaml
```

## Development

Development can be conducted using any GoLang compatible IDE or Text-Editor (e.g., Jetbrains GoLand, VSCode with Go plugins).

### Common tasks

* Build the binary: `go build -ldflags '-linkmode=external' -v -o wget-test-service`
* Run tests: `go test -race -v ./...`
* Build the docker image: `docker build . -t christiankreuzbergerdtx/wget-test-service:dev` (Note: Replace `christiankreuzbergerdtx` with your DockerHub account/organization)
* Push the docker image to DockerHub: `docker push christiankreuzbergerdtx/wget-test-service:dev` (Note: Replace `christiankreuzbergerdtx` with your DockerHub account/organization)
* Deploy the service using `kubectl`: `kubectl apply -f deploy/`
* Undeploy the service using `kubectl`: `kubectl deploy -f deploy/`
* Watch the deployment using `kubectl`: `kubectl -n keptn get deployment wget-test-service -o wide`
* Get logs using `kubectl`: `kubectl -n keptn logs deployment/wget-test-service -f`
* Watch the deployed pods using `kubectl`: `kubectl -n keptn get pods -l run=wget-test-service`
* Deploy the service using [Skaffold](https://skaffold.dev/): `skaffold run --tail` (Note: please adapt the image name in [skaffold.yaml](skaffold.yaml))

### Testing Cloud Events

We have dummy cloud-events in the form of PostMan Requests in the [test-events/](test-events/) directory.

## License

Please find more information in the [LICENSE](LICENSE) file.