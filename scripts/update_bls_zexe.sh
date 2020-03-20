#!/bin/bash -e

BRANCH=${1:-master}
SCRIPTS_DIR=`dirname $0`
cd $SCRIPTS_DIR/..
#rm -r ./external/bls-zexe
#git commit -am"bls-zexe: removes old version"
git subtree add --prefix external/bls-zexe https://github.com/celo-org/bls-zexe $BRANCH --squash
