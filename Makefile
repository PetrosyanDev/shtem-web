# HRACH_DEV © iMed Cloud Services, Inc.
RELEASE_VERSION=1.0.0
ROLLBACK_VERSION=1.0.0

GIT_REPO=git@github.com:PetrosyanDev/shtem-web.git
GIT_BLD_BRANCH=main
DEV_HOST=erik@165.227.148.29
PRD_HOST=erik@165.227.148.29
DEV_BASE=/home/erik
PRD_BASE=/home/erik
SSH_PORT=22
DEPLOY_DIR=shtem-web
# CERTS_DIR:=$(DEV_BASE)/imedcs-tls

IMAGE=shtem-web

## Clonning Repo to Stage Server
init:
	ssh ${DEV_HOST} -p ${SSH_PORT} "git clone ${GIT_REPO}"
	ssh ${DEV_HOST} -p ${SSH_PORT} "cd ${DEPLOY_DIR} && git checkout ${GIT_BLD_BRANCH}"

## Keep Sources UpToDate
pull:
	ssh ${DEV_HOST} -p ${SSH_PORT} "cd ${DEPLOY_DIR} && git checkout ${GIT_BLD_BRANCH} && git pull origin ${GIT_BLD_BRANCH}"

## Generating Protobuff Files
# proto:
# 	@mkdir -p sources/pkg
# 	protoc -I=imedcs-idls/imedcs-api --go_out=sources/pkg --go-grpc_out=sources/pkg imedcs-idls/imedcs-api/api.proto
# 	protoc -I=imedcs-idls/imedcs-storage --go_out=sources/pkg --go-grpc_out=sources/pkg imedcs-idls/imedcs-storage/storage.proto

## Installing all Dependencies on Local Machine
deps:
	npm install
	go mod tidy -compat=1.21.3
	go mod vendor

## Running Tests on Local Machine
test:
	go test -cover shtem-web/...

## Compiling SASS
sass:
	npm run sass

## Compiling TS for Vue Components
vue-compile:
	npm run build
	npm run js:main

## Copy Bootstrap JS lib (anly do on each butstrap versin update, then manualy remove last line for Map source)
js-bootstrap:
	npm run js:bootstrap

## Minifying Resources
minify:
	npm run minify:css

## Running on Local Machine
run: sass vue-compile js-bootstrap minify test
	go run shtem-web/...

## Running on Local Machine with TLS
# run-tls: sass vue-compile js-bootstrap minify test
# 	@NONSENCE=${NONSENCE} go run shtem-web/... --tls --cfg secrets/local.json

build: sass vue-compile js-bootstrap minify test pull
	mkdir -p build/web
	cd sources/cmd/web && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../../build/web/app
	scp -P ${SSH_PORT} -r build ${DEV_HOST}:${DEV_BASE}/${DEPLOY_DIR}/
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} docker-compose -f ${DEPLOY_DIR}/docker/build.yml build"
	@echo "BUILT IMAGE: ${IMAGE}:${RELEASE_VERSION}"

## Building and Deploying on Staging
deploy-dev: 
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} DIR=${DEV_BASE}/${DEPLOY_DIR} MODE=debug NONS=${NONSENCE} docker stack deploy -c ${DEPLOY_DIR}/docker/run.yml erik --with-registry-auth"
	@echo "DEPLOYED on STAGING! VERSION is: ${RELEASE_VERSION}"

deploy-dev-unbuild:
	ssh ${DEV_HOST} -p ${SSH_PORT} "IMG=${IMAGE} TAG=${RELEASE_VERSION} MODE=debug docker stack deploy -c ${DEPLOY_DIR}/docker/run.yml erik --with-registry-auth"
	@echo "DEPLOYED on STAGING! VERSION is: ${RELEASE_VERSION}"

## Deploying on Production (without build)
# deploy-prd: build
# 	ssh ${PRD_HOST} -p ${SSH_PORT} "mkdir -p ${DEPLOY_DIR}/docker"
# 	scp -P ${SSH_PORT} -r docker/run.yml ${PRD_HOST}:${PRD_BASE}/${DEPLOY_DIR}/docker/
# 	scp -P ${SSH_PORT} secrets/prd.json ${PRD_HOST}:${PRD_BASE}/${DEPLOY_DIR}/secrets.json
# 	ssh ${PRD_HOST} -p ${SSH_PORT} "docker pull ${REGISTRY}/${IMAGE}:${RELEASE_VERSION}"
# 	ssh ${PRD_HOST} -p ${SSH_PORT} "REPO=${REGISTRY} IMG=${IMAGE} TAG=${RELEASE_VERSION} DIR=${PRD_BASE}/${DEPLOY_DIR} MODE=release NONS=${NONSENCE} docker stack deploy -c ${DEPLOY_DIR}/docker/run.yml erik --with-registry-auth"
# 	ssh ${PRD_HOST} -p ${SSH_PORT} "rm -f ${DEPLOY_DIR}/secrets.json"
# 	@echo "DEPLOYED on PRODUCTION! VERSION is: ${RELEASE_VERSION}"

## Rolling Back on Production by one deploy
revert:
	ssh ${PRD_HOST} -p ${SSH_PORT} "docker service update --image ${IMAGE}:${ROLLBACK_VERSION} --force erik_${IMAGE} --with-registry-auth"
	@echo "ROLLED BACK on PRODUCTION! now VERSION is: ${ROLLBACK_VERSION}"

## Purge Docker Caches on Build Server
cleanup:
	ssh ${DEV_HOST} -p ${SSH_PORT} "docker system prune -f"