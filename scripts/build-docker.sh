#! /usr/bin/env bash

bail() {
	echo "$1"
	exit 1
}

echo "Checking refname..."

REFNAME=$(git symbolic-ref -q HEAD || git name-rev --no-undefined --name-only --always HEAD)

echo "Refname $REFNAME, validating semver..."

SEMVER=$(echo "$REFNAME" | grep -oE "[0-9]+[.][0-9]+[.][0-9]+")

[ -z $SEMVER ] && { bail "Semver is not valid, stopping build..."; }

echo "Continuing build..."

DOCKER_REGISTRY="host.docker.internal:5000"
BUILD_TAG="igor:$SEMVER"
IMAGE_FULL_NAME="$DOCKER_REGISTRY/$BUILD_TAG"
IMAGE_EXISTS=$(docker pull "$IMAGE_FULL_NAME" > /dev/null 2>&1 && echo "exists")

[ -n $IMAGE_EXISTS ] && { bail "Image already exists: $IMAGE_FULL_NAME"; }

# docker build -t $IMAGE_FULL_NAME .

# docker push $IMAGE_FULL_NAME
