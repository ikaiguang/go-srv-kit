override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_CONFIGS_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_CONFIGS_INTERNAL_PROTO := ""
SAAS_CONFIGS_PROTO_FILES := ""
ifneq ($(SAAS_CONFIGS_INTERNAL_PROTO), "")
	SAAS_CONFIGS_PROTO_FILES=$(SAAS_CONFIGS_API_PROTO) $(SAAS_CONFIGS_INTERNAL_PROTO)
else
	SAAS_CONFIGS_PROTO_FILES=$(SAAS_CONFIGS_API_PROTO)
endif
.PHONY: protoc-config-protobuf
# protoc :-->: generate config api protobuf
protoc-config-protobuf:
	@echo "# generate config api protobuf"
	$(call protoc_protobuf,$(SAAS_CONFIGS_PROTO_FILES))
