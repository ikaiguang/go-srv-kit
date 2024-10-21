override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

KIT_TESTDATA_V1_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
KIT_TESTDATA_V1_INTERNAL_PROTO := ""
KIT_TESTDATA_V1_PROTO_FILES := ""
ifneq ($(KIT_TESTDATA_V1_INTERNAL_PROTO), "")
	KIT_TESTDATA_V1_PROTO_FILES=$(KIT_TESTDATA_V1_API_PROTO) $(KIT_TESTDATA_V1_INTERNAL_PROTO)
else
	KIT_TESTDATA_V1_PROTO_FILES=$(KIT_TESTDATA_V1_API_PROTO)
endif
.PHONY: protoc-testdata-v1-protobuf
# protoc :-->: generate testdata service protobuf
protoc-testdata-v1-protobuf:
	@echo "# generate testdata service v1 protobuf"
	$(call protoc_protobuf,$(KIT_TESTDATA_V1_PROTO_FILES))
