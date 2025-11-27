package pkg

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strconv"

	"github.com/masterfuzz/toysort/pkg/kvheap"
)

func ParseLine(line []byte) *kvheap.KeyVal {
	splits := bytes.Fields(line)
	if len(splits) != 2 {
		log.Printf("WARN: failed to parse line, incorrect format %q", line)
		return nil
	}
	v, err := strconv.ParseInt(string(splits[1]), 10, 64)
	if err != nil {
		panic(err)
	}

	return &kvheap.KeyVal{
		Key: bytes.Clone(splits[0]),
		Val: v,
	}
}

func ToySort(r io.Reader, n int) []kvheap.KeyVal {
	top := kvheap.NewKVTopN(n)
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 16*1024*1024)
	scanner.Buffer(buf, 64*1024)

	num := 0
	for scanner.Scan() {
		num++
		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file, line %d: %v", num, err)
		}

		line := scanner.Bytes()

		p := ParseLine(line)
		if p != nil {
			top.Push(*ParseLine(line))
		}
	}
	log.Printf("read %d lines", num)

	return top.TopN()
}
