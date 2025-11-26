package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/masterfuzz/toysort/pkg"
)

func main() {
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

	top := pkg.NewKVTopN(2)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file %v", err)
		}

		line := scanner.Text()

		top.Push(parseLine(line))
	}

	sorted := top.TopN()
	for _, s := range sorted {
		fmt.Println(s)
	}

}

func parseLine(line string) pkg.KeyVal {
	splits := strings.Split(line, " ")
	v, err := strconv.ParseInt(splits[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return pkg.KeyVal{
		Key: splits[0],
		Val: v,
	}
}
