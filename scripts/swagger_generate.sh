#!/bin/bash -e

SCRIPTS_DIR=`dirname $0`
source $SCRIPTS_DIR/env.sh

openapi-generator generate -g go-server \
    --input-spec "$ROSETTA_DIR/swagger.json" \
    --git-user-id $GITHUB_ORG --git-repo-id $GITHUB_REPO \
    --package-name "api" --additional-properties "sourceFolder=api" \
    --template-dir $ROSETTA_DIR/templates
