#!/bin/bash
HUB="hub-linux-amd64-2.6.0"
CURRENT_DIR=$(pwd)
curl -O -L https://github.com/github/hub/releases/download/v2.6.0/${HUB}.tgz
tar -xzvf ${HUB}.tgz
export PATH=${PATH}:${CURRENT_DIR}/${HUB}/bin
export WINDOWS_SHA=$(sha256sum bin/windows/rbac-view | awk '{ print $1 }' )
export LINUX_SHA=$(sha256sum bin/linux/rbac-view | awk '{ print $1 }' )
export DARWIN_SHA=$(sha256sum bin/darwin/rbac-view | awk '{ print $1 }' )
hub clone https://github.com/jasonrichardsmith/krew-index.git
hub rbac-view.krew.template.yaml | envsubst > krew-index/plugins/rbac-view.yaml
cd krew-index
hub checkout -b ${TRAVIS_TAG}
hub add plugins/rbac-view.yaml
hub commit -m 'Release ${TRAVIS_TAG}'
hub push origin ${TRAVIS_TAG} --force
hub pull-request \
	--base="GoogleContainerTools:master" \
	--head="jasonrichardsmith/krew-index:${TRAVIS_TAG}" \
	--message="Update rbac-view ${TRAVIS_TAG}"
cd ../
rm -rf krew-index ${HUB} ${HUB}.tgz
