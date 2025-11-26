package pkg

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/masterfuzz/toysort/pkg/kvheap"
)

func ParseLine(line string) kvheap.KeyVal {
	splits := strings.Split(line, " ")
	v, err := strconv.ParseInt(splits[1], 10, 64)
	if err != nil {
		panic(err)
	}

	return kvheap.KeyVal{
		Key: splits[0],
		Val: v,
	}
}

func ToySort(r io.Reader, n int) []kvheap.KeyVal {
	top := kvheap.NewKVTopN(n)
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 16*1024*1024)
	scanner.Buffer(buf, 64*1024)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading file %v", err)
		}

		line := scanner.Text()

		top.Push(ParseLine(line))
	}

	return top.TopN()
}
