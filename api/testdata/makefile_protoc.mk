# config
KitTestdata_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/config -name "*.proto")
#KitTestdata_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KitTestdata_INTERNAL_PROTO=
KitTestdata_PROTO_FILES=""
ifneq ($(KitTestdata_INTERNAL_PROTO), "")
	KitTestdata_PROTO_FILES=$(KitTestdata_API_PROTO) $(KitTestdata_INTERNAL_PROTO)
else
	KitTestdata_PROTO_FILES=$(KitTestdata_API_PROTO)
endif
.PHONY: protoc-testdata-protobuf
# protoc :-->: generate testdata protobuf
protoc-testdata-protobuf:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(KitTestdata_PROTO_FILES))
