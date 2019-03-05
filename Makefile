PATH := $(CURDIR)/bin:$(PATH)

protoc: tools
	@$(MAKE) -C api/proto all

TOOLS += github.com/golang/protobuf/protoc-gen-go
tools: $(TOOLS)

$(TOOLS): %:
	GOBIN=$(CURDIR)/bin go install $*

images:
	docker build -f build/client.store.svc/Dockerfile -t 90poe/client.store.svc:latest .

csv.consumer.cmd:
	go build -o bin/csv.consumer.cmd ./cmd/csv.consumer.cmd

client.store.svc:
	go build -o bin/client.store.svc ./cmd/client.store.svc
