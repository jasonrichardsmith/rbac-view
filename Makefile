DOMAIN="jasonrichardsmith"
APP="rbac-view"
HUB="hub-linux-amd64-2.6.0"
USERID=$(shell id -u)
GROUPID=$(shell id -g)

buildgodocker:
	docker run \
	-v $(PWD):/go/src/github.com/$(DOMAIN)/$(APP) \
	-w /go/src/github.com/$(DOMAIN)/$(APP) \
	--entrypoint '/bin/bash' \
	jasonrichardsmith/glide_builder:2 \
	-c "make buildgo && \
	chown -R $(USERID):$(GROUPID) bin && \
	chown -R $(USERID):$(GROUPID) vendor"

depgodocker:
	docker run \
	-v $(PWD):/go/src/github.com/$(DOMAIN)/$(APP) \
	-w /go/src/github.com/$(DOMAIN)/$(APP) \
	--entrypoint '/bin/bash' \
	jasonrichardsmith/glide_builder:2 \
	-c "make godep && \
	chown -R $(USERID):$(GROUPID) vendor"

buildnpmdocker:
	docker run -v "$(PWD)":/usr/src/app -w /usr/src/app \
	--entrypoint '/bin/bash' \
	node:6 \
	-c "make buildnpm && \
	chown -R $(USERID):$(GROUPID) /usr/src/app/frontend/*"

npmdep:
	cd frontend; npm install
godep: godepfb
	glide install
	go generate
godepfb:
	go get -u github.com/UnnoTed/fileb0x
buildnpm: npmdep
	cd frontend; npm run build
buildgo: godep
	mkdir -p bin/linux
	mkdir -p bin/windows
	mkdir -p bin/darwin
	GOOS=linux go build -a -installsuffix cgo -o bin/linux/rbac-view
	GOOS=windows go build -a -installsuffix cgo -o bin/windows/rbac-view
	GOOS=darwin go build -a -installsuffix cgo -o bin/darwin/rbac-view
distclean:
	rm -rf vendor
	rm -rf frontend/node_modules
clean: distclean
	rm -rf bin
	rm frontend/dist/build.css frontend/dist/build.js

builddocker: buildnpmdocker buildgodocker
build: buildnpm buildgo

releases: 
	tar -czvf rbac-view.$(TAG).linux.tar.gz bin/linux/rbac-view
	tar -czvf rbac-view.$(TAG).windows.tar.gz bin/windows/rbac-view
	tar -czvf rbac-view.$(TAG).darwin.tar.gz bin/darwin/rbac-view

.PHONY: krew-index
krew-index:
	curl -O -L https://github.com/github/hub/releases/download/v2.6.0/$(HUB).tgz
	tar -xzvf $(HUB).tgz
	git clone https://github.com/jasonrichardsmith/krew-index.git
	$(eval CURRENT_DIR=$(shell pwd))
	$(eval export TAG)
	$(eval export WINDOWS_SHA=$(shell sha256sum bin/windows/rbac-view | awk '{ print $$1 }' ))
	$(eval export LINUX_SHA=$(shell sha256sum bin/linux/rbac-view | awk '{ print $$1 }' ))
	$(eval export DARWIN_SHA=$(shell sha256sum bin/darwin/rbac-view | awk '{ print $$1 }' ))
	envsubst < rbac-view.krew.template.yaml > krew-index/plugins/rbac-view.yaml
	cd krew-index && \
		git checkout -b $(TAG) && \
		cat plugins/rbac-view.yaml && \
		git add plugins/rbac-view.yaml && \
		git commit --author="jasonrichardsmith <jasonrichardsmith@gmail.com>" -m 'Release $(TAG)' && \
		git remote add krew-index \
		https://$(GITHUB_TOKEN)@github.com/jasonrichardsmith/krew-index.git > /dev/null 2>&1 && \
		git push --quiet --set-upstream krew-index $(TAG) --force && \
		$(CURRENT_DIR)/$(HUB)/bin/hub pull-request \
			--base="GoogleContainerTools:master" \
			--head="jasonrichardsmith/krew-index:${TAG}" \
			--message="Update rbac-view ${TAG}"

.PHONY: full-release
full-release: | releases krew-index 
