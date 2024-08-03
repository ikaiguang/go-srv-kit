# config
KitPing_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/config -name "*.proto")
#KitPing_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KitPing_INTERNAL_PROTO=
KitPing_PROTO_FILES=""
ifneq ($(KitPing_INTERNAL_PROTO), "")
	KitPing_PROTO_FILES=$(KitPing_API_PROTO) $(KitPing_INTERNAL_PROTO)
else
	KitPing_PROTO_FILES=$(KitPing_API_PROTO)
endif
.PHONY: protoc-ping-protobuf
# protoc :-->: generate ping protobuf
protoc-ping-protobuf:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(KitPing_PROTO_FILES))
