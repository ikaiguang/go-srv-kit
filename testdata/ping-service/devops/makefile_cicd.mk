.PHONY: rm-none-image
# rm none:none images
rm-none-image:
	@echo "rm :-->: rm none images"
	/bin/sh ./testdata/ping-service/devops/docker-build/service-image/rm_none_images.sh

.PHONY: build-base-image
# build :-->: base image
build-base-image:
	@echo "build :-->: build base image"
	/bin/sh ./testdata/ping-service/devops/docker-build/service-image/build_base_image.sh
	/bin/sh ./testdata/ping-service/devops/docker-build/service-image/build_release_image.sh
	$(MAKE) rm-none-image

.PHONY: build-service-image
# build :-->: service image
build-service-image:
	@echo "build :-->: build service image"
	$(MAKE) build-base-image
	/bin/sh ./testdata/ping-service/devops/docker-build/service-image/build_service_image.sh
	$(MAKE) rm-none-image

# build-image
.PHONY: build
# build :-->: service image
build:
	#docker build -t ping-service -f ./testdata/ping-service/devops/Dockerfile .
	#docker pull golang:1.22.8
	#docker pull debian:stable-20240926-slim
	docker build \
		--build-arg BUILD_FROM_IMAGE=golang:1.22.8 \
		--build-arg RUN_SERVICE_IMAGE=debian:stable-20240926-slim \
		--build-arg APP_DIR=testdata \
		--build-arg SERVICE_NAME=ping-service \
		--build-arg VERSION=latest \
		-t ping-service:latest \
		-f ./testdata/ping-service/devops/docker-build/Dockerfile .

.PHONY: deploy-general-config
# deploy :-->: store general config
deploy-general-config:
	@echo "deploy :-->: general config"
	go run ./testdata/ping-service/cmd/store-configuration/... \
      -conf=./testdata/ping-service/configs \
      -source_dir ./testdata/ping-service/devops/docker-deploy/general-configs \
      -store_dir go-micro-saas/general-configs/testing

.PHONY: deploy-service-config
# deploy :-->: store service config
deploy-service-config:
	@echo "deploy :-->: service config"
	go run ./testdata/ping-service/cmd/store-configuration/... \
      -conf=./testdata/ping-service/configs \
      -source_dir ./testdata/ping-service/devops/docker-deploy/service-configs \
      -store_dir go-micro-saas/ping-service/testing/v1.0.0

.PHONY: deploy-database-migration
# deploy :-->: database migration
deploy-database-migration:
	@echo "deploy :-->: database migration"
	$(MAKE) run-database-migration

# deploy-image on docker
.PHONY: deploy-on-docker
# deploy-on-docker :-->: deploying on docker
deploy-on-docker:
	docker-compose -f ./testdata/ping-service/devops/docker-deploy/docker-compose.yaml up -d

.PHONY: stop-docker-deploy
# stop :-->: docker container
stop-docker-deploy:
	@echo "deploy-on-docker :-->: deploying on docker"
	docker-compose -f ./testdata/ping-service/devops/docker-deploy/docker-compose.yaml down

