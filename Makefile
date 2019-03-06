PATH := $(CURDIR)/bin:$(PATH)

DB_HOST            ?= localhost
DB_NAME            ?= client.store.svc
DB_USER            ?= postgres
DB_PORT            ?= 5432
DB_PASSWORD        ?= postgres
DB_SSL_MODE        ?= disable
DB_MIGRATIONS_PATH ?= $(CURDIR)/migrations

protoc: tools
	@$(MAKE) -C api/proto all

TOOLS += github.com/golang/protobuf/protoc-gen-go
TOOLS += github.com/pressly/goose/cmd/goose
tools: $(TOOLS)

$(TOOLS): %:
	GOBIN=$(CURDIR)/bin go install $*

images:
	docker build -f build/client.store.svc/Dockerfile -t bailey/client.store.svc:latest .

csv.consumer.cmd:
	go build -o bin/csv.consumer.cmd ./cmd/csv.consumer.cmd

client.store.svc:
	go build -o bin/client.store.svc ./cmd/client.store.svc

test:
	go test -v ./...

test-integration:
	go test -v -tags integration ./...

PSQL_URI = "host=$(DB_HOST) user=$(DB_USER) dbname=$(DB_NAME) sslmode=$(DB_SSL_MODE) password=$(DB_PASSWORD)"

migrate_up:
	goose -dir migrations postgres $(PSQL_URI) up

migrate_down:
	goose -dir migrations postgres $(PSQL_URI) down

run: images up

up:
	docker-compose -f deployments/docker-compose.yml up -d

down:
	docker-compose -f deployments/docker-compose.yml down
