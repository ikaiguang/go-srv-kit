# saas services
SAAS_SERVICE_PROTO_FILES=$(shell cd $(PROJECT_PATH) && find api -name "*.proto")
.PHONY: protoc-api-protobuf
# protoc :-->: generate api protobuf
protoc-api-protobuf:
	@echo "# generate api protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_PROTO_FILES))

# specified server
SAAS_SERVICE_SPECIFIED_FILES=$(shell cd $(PROJECT_PATH) && find ./api/${service} -name "*.proto")
.PHONY: protoc-specified-api
# protoc :-->: example: make protoc-specified-api service=ping-service
protoc-specified-api:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(SAAS_SERVICE_SPECIFIED_FILES))
