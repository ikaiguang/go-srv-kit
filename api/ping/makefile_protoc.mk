# config
KIT_PING_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/ping -name "*.proto")
#KIT_PING_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KIT_PING_INTERNAL_PROTO=
KIT_PING_PROTO_FILES=""
ifneq ($(KIT_PING_INTERNAL_PROTO), "")
	KIT_PING_PROTO_FILES=$(KIT_PING_API_PROTO) $(KIT_PING_INTERNAL_PROTO)
else
	KIT_PING_PROTO_FILES=$(KIT_PING_API_PROTO)
endif
.PHONY: protoc-ping-protobuf
# protoc :-->: generate ping service protobuf
protoc-ping-protobuf:
	@echo "# generate ping service protobuf"
	$(call protoc_protobuf,$(KIT_PING_PROTO_FILES))
