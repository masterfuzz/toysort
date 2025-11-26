package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"flag"

	"github.com/masterfuzz/toysort/pkg"
)

func main() {
	topN := flag.Int("n", 10, "maximum entries in the output")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	var fname string
	if scanner.Scan() {
		fname = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	sorted := pkg.ToySort(file, *topN)
	for _, s := range sorted {
		fmt.Println(s.Key)
	}

}
