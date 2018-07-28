DOMAIN="jasonrichardsmith"
APP="rbac-view"
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
