OUT    = ../../
PROTOS = $(shell find . -type f -name '*.proto')

HAS_PROTOC := $(shell command -v protoc;)
ifndef HAS_PROTOC
	$(error Install protoc: https://github.com/protocolbuffers/protobuf/releases)
endif

all: $(PROTOS)

.PHONY: $(PROTOS)
$(PROTOS): %.proto:
	$(info Generating $*.proto)
	@protoc -I . --go_out=plugins=grpc:$(OUT) $*.proto
