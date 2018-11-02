#!/bin/bash
HUB="hub-linux-amd64-2.6.0"
CURRENT_DIR=$(pwd)
export TAG="${TRAVIS_TAG:-test-tag}"
curl -O -L https://github.com/github/hub/releases/download/v2.6.0/${HUB}.tgz
tar -xzvf ${HUB}.tgz
export PATH=${PATH}:${CURRENT_DIR}/${HUB}/bin
export WINDOWS_SHA=$(sha256sum bin/windows/rbac-view | awk '{ print $1 }' )
export LINUX_SHA=$(sha256sum bin/linux/rbac-view | awk '{ print $1 }' )
export DARWIN_SHA=$(sha256sum bin/darwin/rbac-view | awk '{ print $1 }' )
git clone https://github.com/jasonrichardsmith/krew-index.git
envsubst < rbac-view.krew.template.yaml > krew-index/plugins/rbac-view.yaml
cd krew-index
git checkout -b ${TAG}
git add plugins/rbac-view.yaml
git commit -m 'Release ${TAG}'
git remote add krew-index https://${GITHUB_TOKEN}@github.com/jasonrichardsmith/krew-index.git > /dev/null 2>&1
git push --quiet --set-upstream krew-index ${TAG} --force

#hub pull-request \
#	--base="GoogleContainerTools:master" \
#	--head="jasonrichardsmith/krew-index:${TAG}" \
#	--message="Update rbac-view ${TAG}"

cd ../
rm -rf krew-index ${HUB} ${HUB}.tgz
