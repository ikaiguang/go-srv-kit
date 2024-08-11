# config
KIT_TESTDATA_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/testdata -name "*.proto")
#KIT_TESTDATA_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KIT_TESTDATA_INTERNAL_PROTO=
KIT_TESTDATA_PROTO_FILES=""
ifneq ($(KIT_TESTDATA_INTERNAL_PROTO), "")
	KIT_TESTDATA_PROTO_FILES=$(KIT_TESTDATA_API_PROTO) $(KIT_TESTDATA_INTERNAL_PROTO)
else
	KIT_TESTDATA_PROTO_FILES=$(KIT_TESTDATA_API_PROTO)
endif
.PHONY: protoc-testdata-protobuf
# protoc :-->: generate testdata service protobuf
protoc-testdata-protobuf:
	@echo "# generate testdata service protobuf"
	$(call protoc_protobuf,$(KIT_TESTDATA_PROTO_FILES))
