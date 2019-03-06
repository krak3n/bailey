package rpc

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/krak3n/bailey/internal/storage"
	"github.com/krak3n/bailey/pkg/api/clientstoresvc"
)

// The ClientStoreService implements the proto API for the client store servive
type ClientStoreService struct {
	store storage.Store
}

// NewClientStoreService constructs a new client store service
func NewClientStoreService(store storage.Store) *ClientStoreService {
	return &ClientStoreService{
		store: store,
	}
}

// Upsert handles a stream of client upsert requests, processing each one at
// time by calling the underlying storage layer
func (svc *ClientStoreService) Upsert(stream clientstoresvc.ClientStoreService_UpsertServer) error {
	var processed int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err // TODO: Status codes with status details
		}
		// TODO: format phone numbers
		log.Println(fmt.Sprintf("upserting client: %d", req.Client.Id))
		err = svc.store.UpsertClient(context.Background(), &storage.Client{
			ID:          req.Client.Id,
			Name:        req.Client.Name,
			PhoneNumber: req.Client.PhoneNumber,
			Email:       req.Client.Email,
		})
		if err != nil {
			return err
		}
		processed++
	}

	return stream.SendAndClose(&clientstoresvc.UpsertResponse{
		Processed: processed,
	})
}
