package psql

import (
	"context"

	"github.com/krak3n/bailey/internal/storage"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Store holds methods for performing operations againsts a PSQL database
type Store struct {
	session sqlbuilder.Database
}

// New constructs a new Store
func New(session sqlbuilder.Database) *Store {
	return &Store{
		session: session,
	}
}

// UpsertClient inserts or updates a client row within a single DB transaction
func (s *Store) UpsertClient(ctx context.Context, client *storage.Client) error {
	return s.session.Tx(ctx, func(tx sqlbuilder.Tx) error {
		// TODO: consider a Insert On Conflict Update, avoids a read first
		res := tx.Collection(storage.ClientsTable).Find(db.Cond{"id": client.ID})
		total, err := res.Count()
		if err != nil {
			return err
		}
		if total == 0 {
			return tx.Collection(storage.ClientsTable).InsertReturning(client)
		}
		return res.Update(client)
	})
}
