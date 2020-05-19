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

LAST_VERSION=$(grep "MiddlewareVersion" $VERSION_FILE | sed -e 's/.*\= "//' -e 's/\".*//')

echo "Current version: $LAST_VERSION"
echo "What's the next version?"
read NEW_VERSION

sed -i -e "s/$LAST_VERSION/$NEW_VERSION/" $VERSION_FILE

git checkout -b "release/${NEW_VERSION}"
git add $VERSION_FILE
git ci -m "Release ${NEW_VERSION}"
git tag "v${NEW_VERSION}"

COMMIT_SHA=$(git rev-parse HEAD)

echo "Building Docker images"
docker build \
  --build-arg COMMIT_SHA=${COMMIT_SHA} \
  -t us.gcr.io/celo-testnet/rosetta:${COMMIT_SHA} \
  -t us.gcr.io/celo-testnet/rosetta:${NEW_VERSION} \
  -t us.gcr.io/celo-testnet/rosetta:latest \
  .

echo "Pushing Docker images"
docker push us.gcr.io/celo-testnet/rosetta:${COMMIT_SHA} 
docker push us.gcr.io/celo-testnet/rosetta:${NEW_VERSION} 
docker push us.gcr.io/celo-testnet/rosetta:latest 

echo "Pushing to git"
git push origin "release/${NEW_VERSION}"
git push origin "v${NEW_VERSION}"
