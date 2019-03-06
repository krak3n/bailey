package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var csvpath = flag.String("csv", "", "CSV file to process")

var closeC = make(chan bool, 0)

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

	r := csv.NewReader(f)

	var wg sync.WaitGroup
	wg.Add(2)

	linesC := make(chan []string, 1)
	go func() {
		defer wg.Done()
		readLines(r, (chan<- []string)(linesC))
	}()

	go func() {
		defer wg.Done()
		processLines((<-chan []string)(linesC))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println(<-sigC)

	close(closeC)

	wg.Wait()
	check(f.Close())
}

// readLines reads lines from a reader and puts them on a channel for processing
func readLines(r reader, c chan<- []string) {
	for {
		select {
		case <-closeC:
			return
		default:
			line, err := r.Read()
			if err == io.EOF {
				fmt.Println("Done")
				return
			}
			check(err)
			c <- line
		}
	}
}

// processLines takes channel of csv lines to process
// TODO: store state of where got to in processing the CSV?
func processLines(c <-chan []string) {
	var i int
	for {
		select {
		case <-closeC:
			return
		case line := <-c:
			i++
			if i == 1 { // skip header
				continue
			}
			fmt.Println(line)
		}
	}
}
