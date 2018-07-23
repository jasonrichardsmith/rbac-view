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
	chown $(USERID):$(GROUPID) rbac_view && \
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
	go build -a -installsuffix cgo -o rbac_view
distclean:
	rm -rf vendor
	rm -rf frontend/node_modules
clean: distclean
	rm -rf rbac_view
	rm frontend/dist/build.css frontend/dist/build.js

