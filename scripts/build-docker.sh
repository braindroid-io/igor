#! /usr/bin/env bash

echo "Checking refname..."

REFNAME=$(git symbolic-ref -q HEAD || git name-rev --no-undefined --name-only --always HEAD)

echo "Refname $REFNAME, validating semver..."

SEMVER=$(echo "$REFNAME" | grep -oE "[0-9]+[.][0-9]+[.][0-9]+")

[ -z $SEMVER ] && { echo "Semver is not valid, stopping build..."; exit 1; }

DOCKER_REGISTRY="host.docker.internal:5000"
BUILD_TAG="igor:$SEMVER"
IMAGE_FULL_NAME="$DOCKER_REGISTRY/$BUILD_TAG"
IMAGE_EXISTS=$(docker pull "$IMAGE_FULL_NAME" > /dev/null 2>&1 && echo "exists")

[ -n $IMAGE_EXISTS ] && { echo "Image already exists: $IMAGE_FULL_NAME"; exit 1; }

# docker build -t $BUILD_TAG .
#
# docker tag $BUILD_TAG $IMAGE_FULL_NAME
#
# docker push $IMAGE_FULL_NAME
