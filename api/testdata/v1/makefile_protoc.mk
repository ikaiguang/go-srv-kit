# config
KIT_TESTDATA_V1_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/testdata/v1 -name "*.proto")
#KIT_TESTDATA_V1_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
KIT_TESTDATA_V1_INTERNAL_PROTO=
KIT_TESTDATA_V1_PROTO_FILES=""
ifneq ($(KIT_TESTDATA_V1_INTERNAL_PROTO), "")
	KIT_TESTDATA_V1_PROTO_FILES=$(KIT_TESTDATA_V1_API_PROTO) $(KIT_TESTDATA_V1_INTERNAL_PROTO)
else
	KIT_TESTDATA_V1_PROTO_FILES=$(KIT_TESTDATA_V1_API_PROTO)
endif
.PHONY: protoc-testdata-v1-protobuf
# protoc :-->: generate testdata protobuf
protoc-testdata-v1-protobuf:
	@echo "# generate ${service} protobuf"
	$(call protoc_protobuf,$(KIT_TESTDATA_V1_PROTO_FILES))
