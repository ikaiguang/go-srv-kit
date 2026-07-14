override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

# saas services
.PHONY: protoc-api-protobuf
# protoc :-->: generate services api protobuf
protoc-api-protobuf:
	@echo "# generate services api protobuf"
	$(MAKE) protoc-ping-protobuf

# specified server
SAAS_SERVICE_SPECIFIED_FILES := $(shell find ./$(REL_PROJECT_PATH)/${service} -name "*.proto")
.PHONY: protoc-specified-api
# protoc :-->: example: make protoc-specified-api service=ping-service
protoc-specified-api:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_SPECIFIED_FILES))

