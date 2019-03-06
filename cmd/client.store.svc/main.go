package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/krak3n/bailey/internal/rpc"
	"github.com/krak3n/bailey/internal/storage"
	"github.com/krak3n/bailey/internal/storage/psql"
	"github.com/krak3n/bailey/pkg/api/clientstoresvc"
	"google.golang.org/grpc"
	"upper.io/db.v3/postgresql"
)

var bind = flag.String("bind", "0.0.0.0:3000", "server bind address")

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *bind)
	check(err)

	db, err := store()
	check(err)

	var wg sync.WaitGroup
	wg.Add(1)

	srv := grpc.NewServer()
	clientstoresvc.RegisterClientStoreServiceServer(srv, rpc.NewClientStoreService(db))
	go func() {
		defer wg.Done()
		log.Println("starting grpc server:", *bind)
		if err := srv.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.Println(<-sigC)

	log.Println("stopping server")
	srv.Stop() // TODO: graceful stop - the handler will need a Close method

	wg.Wait()
}

// store constructs a new storage layer
func store() (storage.Store, error) {
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
	if err != nil {
		return nil, err
	}
	return psql.New(session), nil
}
