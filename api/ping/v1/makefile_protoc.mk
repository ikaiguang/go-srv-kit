# config
KIT_PING_V1_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/ping/v1 -name "*.proto")
#KIT_PING_V1_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KIT_PING_V1_INTERNAL_PROTO=
KIT_PING_V1_PROTO_FILES=""
ifneq ($(KIT_PING_V1_INTERNAL_PROTO), "")
	KIT_PING_V1_PROTO_FILES=$(KIT_PING_V1_API_PROTO) $(KIT_PING_V1_INTERNAL_PROTO)
else
	KIT_PING_V1_PROTO_FILES=$(KIT_PING_V1_API_PROTO)
endif
.PHONY: protoc-ping-v1-protobuf
# protoc :-->: generate ping protobuf
protoc-ping-v1-protobuf:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(KIT_PING_V1_PROTO_FILES))
