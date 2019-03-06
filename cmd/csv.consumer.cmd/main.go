package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var csvpath = flag.String("csv", "", "CSV file to process") // nolint: gochecknoglobals

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
	for i := 0; ; i++ {
		line, err := r.Read()
		if err == io.EOF {
			fmt.Println("Done")
		}
		check(err)
		if i == 0 { // Skip the header
			continue
		}
		fmt.Println(line)
	}

}
