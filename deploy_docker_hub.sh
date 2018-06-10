#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker build --tag "yanc0/kubernetes-updown-io-controller:$TRAVIS_TAG" .
docker push "yanc0/kubernetes-updown-io-controller:$TRAVIS_TAG"