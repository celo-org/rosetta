#!/bin/bash

set -e 

VERSION_FILE="service/rpc/versions.go"

echo "Fetching last changes from git"
git fetch

if [ -n "$(git status --porcelain)" ]; then 
  echo "There are uncommitted changes. Can't proceed"
  exit 1
fi

if [[ $(git rev-parse --abbrev-ref HEAD) != "master" ]]; then 
  echo "Not in master branch. Can't proceed"
  exit 2
fi

if [[ $(git rev-parse origin/master) != $(git rev-parse HEAD) ]]; then
  echo "Branch not up-to-date. Can't proceed"
  exit 3
fi

VERSION=$(grep "MiddlewareVersion" $VERSION_FILE | sed -e 's/.*\= "//' -e 's/\".*//')

COMMIT_SHA=$(git rev-parse HEAD)

echo "Building Docker images"
docker build \
  --build-arg COMMIT_SHA=${COMMIT_SHA} \
  -t us.gcr.io/celo-testnet/rosetta:${COMMIT_SHA} \
  -t us.gcr.io/celo-testnet/rosetta:${VERSION} \
  -t us.gcr.io/celo-testnet/rosetta:latest \
  .

# We tag the version after building, so that we don't get stuck due to existing
# tags, if the build fails and we need to re-run this script.
echo "Tagging commit with v${VERSION}"
git tag "v${VERSION}"
git push origin "v${VERSION}"

echo "Pushing Docker images"
docker push us.gcr.io/celo-testnet/rosetta:${COMMIT_SHA} 
docker push us.gcr.io/celo-testnet/rosetta:${VERSION} 
docker push us.gcr.io/celo-testnet/rosetta:latest 


echo "Tagged commit ${COMMIT_SHA} with v${VERSION}"
echo "Pushed the following docker images:"
echo "push us.gcr.io/celo-testnet/rosetta:${COMMIT_SHA}"
echo "push us.gcr.io/celo-testnet/rosetta:${VERSION}"
echo "push us.gcr.io/celo-testnet/rosetta:latest"
