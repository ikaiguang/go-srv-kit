override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

KIT_PING_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
KIT_PING_INTERNAL_PROTO := ""
KIT_PING_PROTO_FILES = ""
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
