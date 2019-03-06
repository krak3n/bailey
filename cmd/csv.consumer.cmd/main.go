package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var csvpath = flag.String("csv", "", "CSV file to process")

var sigC = make(chan os.Signal, 1)

type reader interface {
	Read() ([]string, error)
}

// Checks errors, should us a proper logger
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	f, err := os.Open(*csvpath)
	check(err)

	r := csv.NewReader(bufio.NewReader(f))
	go readCSV(r)

	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-sigC
	fmt.Println(sig)
}

func readCSV(r reader) {
	for i := 0; ; i++ {
		line, err := r.Read()
		if err == io.EOF {
			fmt.Println("Done")
			break
		}
		check(err)
		if i == 0 { // Skip the header
			continue
		}
		fmt.Println(line)
	}
}
