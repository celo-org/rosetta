#!/bin/bash -e

SCRIPTS_DIR=`dirname $0`
source $SCRIPTS_DIR/env.sh

openapi-generator generate -g go-server -i $SCRIPTS_DIR/../swagger.json --additional-properties="packageName=api,sourceFolder=api" --git-user-id=$GITHUB_ORG --git-repo-id=$GITHUB_REPO
