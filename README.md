# Kubernetes updown.io controller

**Important**: this project was created for example purpose only. It is not intended to
be ran on any important environment.

[![Build Status](https://travis-ci.org/yanc0/kubernetes-updown-io-controller.svg?branch=master)](https://travis-ci.org/yanc0/kubernetes-updown-io-controller)

## Usage

### Fill you updown.io API key in controller's configMap

`vim ./manifests/kubernetes-updown-io-controller/configmap.yaml`

### Install kubernetes-updown-io-controller

`kubectl apply -f ./manifests/kubernetes-updown-io-controller`

### Add new checks

`kubectl apply -f ./manifests/examples`

## Devel

* Your Go environment is properly configured
* Glide (vendor management for go) is installed. Dep isn't working well with kubernetes go client yet.

### Compile

```
$ mkdir -p $GOPATH/src/github.com/yanc0/
$ cd $GOPATH/src/github.com/yanc0/ && git clone git@github.com:yanc0/kubernetes-updown-io-controller
$ glide up -v
$ go build
```

### Launch

* Get you super powered updown.io api key [https://updown.io/settings/edit](https://updown.io/settings/edit)

`./kubernetes-updown-io-controller -kubeconfig=~/.kube/config -apikey=XXXXXXXXXXXX`