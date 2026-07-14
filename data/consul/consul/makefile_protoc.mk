override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

CONFIG_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
CONFIG_INTERNAL_PROTO := ""
CONFIG_PROTO_FILES := ""
ifneq ($(CONFIG_INTERNAL_PROTO), "")
	CONFIG_PROTO_FILES=$(CONFIG_PROTO) $(CONFIG_INTERNAL_PROTO)
else
	CONFIG_PROTO_FILES=$(CONFIG_PROTO)
endif
.PHONY: protoc-config-protobuf
# protoc :-->: generate config protobuf
protoc-config-protobuf:
	@echo "# generate config protobuf"
	$(call protoc_protobuf,$(CONFIG_PROTO_FILES))
