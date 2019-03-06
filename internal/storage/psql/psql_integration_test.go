// +build integration

package psql

import (
	"context"
	"os"
	"testing"

	"github.com/krak3n/bailey/internal/storage"
	"github.com/stretchr/testify/require"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

func TestUpsertClient(t *testing.T) {
	ctx := context.Background()

	settings := postgresql.ConnectionURL{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		Options: map[string]string{
			"sslmode": os.Getenv("DB_SSL_MODE"),
		},
	}

	session, err := postgresql.Open(settings)
	require.NoError(t, err)

	store := New(session)
	client := &storage.Client{
		ID:          1,
		Name:        "Foo",
		PhoneNumber: "12345678",
		Email:       "foo@bar.com",
	}

	require.NoError(t, store.UpsertClient(ctx, client))

	var c = new(storage.Client)
	err = session.Tx(ctx, func(tx sqlbuilder.Tx) error {
		return tx.Collection(storage.ClientsTable).
			Find(db.Cond{"id": client.ID}).
			One(c)
	})
	require.NoError(t, err)

	require.Equal(t, client, c)
}
