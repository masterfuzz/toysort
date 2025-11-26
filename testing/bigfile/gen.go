package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"

	"github.com/masterfuzz/toysort/pkg/kvheap"
)

func main() {
	lines := flag.Int("l", 16000000, "number of lines to generate (not including the answers)")
	topN := flag.Int("n", 10, "how many answers to generate")
	questionFile := flag.String("q", "./question.txt", "question file location")
	answerFile := flag.String("a", "./answer.txt", "answer file location")
	flag.Parse()

	// create a large question file (~865MB) and the expected top 10 in an answer file

	log.Println("generating test files")
	form := "https://api.tech.com/item/%d"
	var maxInt64 int64 = math.MaxInt64
	var halfMax int64 = math.MaxInt64 / 2

	// first create the expected top 10
	ans := []kvheap.KeyVal{}
	for i := range *topN {
		ans = append(ans, kvheap.KeyVal{
			Key: fmt.Sprintf(form, i),
			Val: halfMax + 1 + rand.Int63n(maxInt64-halfMax),
		})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Val > ans[j].Val
	})

	qf, err := os.Create(*questionFile)
	if err != nil {
		panic(err)
	}
	defer qf.Close()
	bq := bufio.NewWriter(qf)

	af, err := os.Create(*answerFile)
	if err != nil {
		panic(err)
	}
	defer af.Close()

	for _, kv := range ans {
		fmt.Fprintf(af, "%s\n", kv.Key)
		fmt.Fprintf(bq, "%s %d\n", kv.Key, kv.Val)
	}
	af.Close()

	for i := range *lines {
		kv := kvheap.KeyVal{
			Key: fmt.Sprintf(form, i+11),
			Val: rand.Int63n(halfMax),
		}
		fmt.Fprintf(bq, "%s %d\n", kv.Key, kv.Val)
	}

	log.Printf("generated %d lines", *lines+*topN)
}
