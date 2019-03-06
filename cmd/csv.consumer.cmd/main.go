package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/krak3n/bailey/internal/csv"
	"github.com/krak3n/bailey/pkg/api/clientstoresvc"
	"google.golang.org/grpc"
)

var (
	csvpath = flag.String("csv", "", "CSV file to process")
	server  = flag.String("", "127.0.0.1:3000", "server address")
	closeC  = make(chan bool, 0)
)

type reader interface {
	Read() ([]string, error)
}

// Checks errors, should us a proper logger
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: this code certainly needs a cleanup
func main() {
	flag.Parse()

	f, err := os.Open(*csvpath)
	check(err)

	syncer := csv.NewSyncer(f)

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	check(err)

	client := clientstoresvc.NewClientStoreServiceClient(conn)
	stream, err := client.Upsert(context.Background())
	check(err)

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		if err := syncer.SyncTo(write(stream)); err != nil {
			log.Println(err)
		}
		sigC <- syscall.SIGQUIT
	}()

	log.Println(<-sigC)
	syncer.Close()

	rsp, err := stream.CloseAndRecv()
	switch err {
	case nil:
		log.Println(fmt.Sprintf("Processed: %d", rsp.Processed))
	default:
		log.Println(err)
	}

	if err := conn.Close(); err != nil {
		log.Println(err)
	}

	if err := f.Close(); err != nil {
		log.Println(err)
	}
}

func write(stream clientstoresvc.ClientStoreService_UpsertClient) csv.WriteFunc {
	return func(record []string) error {
		id, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return err
		}
		return stream.Send(&clientstoresvc.UpsertRequest{
			Client: &clientstoresvc.Client{
				Id:          id,
				Name:        record[1],
				Email:       record[2],
				PhoneNumber: record[3], // TODO: phone formatting
			},
		})
	}
}
