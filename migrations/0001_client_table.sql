-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE clients (
	"id" integer PRIMARY KEY,
	"name" text NOT NULL,
	"phone_number" text NOT NULL,
	"email" text NOT NULL
) WITH (OIDS = FALSE);

-- +goose Down
-- SQL in this section is executed when the migration is applied.

DROP TABLE clients;
