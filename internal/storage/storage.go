package storage

import "context"

// ClientsTable is the name of the clients database table
var ClientsTable = "clients"

// A Store can do database operations
type Store interface {
	ClientUpserter
}

// A ClientUpserter can upsert a client row
type ClientUpserter interface {
	UpsertClient(context.Context, *Client) error
}

// Client represents a client row
type Client struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
}
