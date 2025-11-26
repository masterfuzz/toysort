package pkg

import (
	"bytes"
	"fmt"
	"sort"
	"testing"

	"github.com/masterfuzz/toysort/pkg/kvheap"
)

// basic fuzz test example courtesy of Claude
func FuzzToySort(f *testing.F) {
	// Seed corpus with some examples
	f.Add(3, []byte("a"), int64(10), []byte("b"), int64(20), []byte("c"), int64(5))
	f.Add(2, []byte("foo"), int64(100), []byte("bar"), int64(50), []byte(""), int64(0))
	f.Add(1, []byte("x"), int64(42), []byte(""), int64(0), []byte(""), int64(0))

	f.Fuzz(func(t *testing.T, n int, key1 []byte, val1 int64, key2 []byte, val2 int64, key3 []byte, val3 int64) {
		// Skip invalid inputs
		if n <= 0 {
			t.Skip()
		}

		// Build input and expected items from fuzz parameters
		var inputBuilder bytes.Buffer
		var allItems []kvheap.KeyVal

		// Helper to add a key-value pair
		addKeyVal := func(key []byte, val int64) {
			// Keys must be non-whitespace
			keyStr := string(bytes.ReplaceAll(key, []byte{' '}, []byte{}))
			keyStr = string(bytes.ReplaceAll([]byte(keyStr), []byte{'\n'}, []byte{}))
			keyStr = string(bytes.ReplaceAll([]byte(keyStr), []byte{'\t'}, []byte{}))
			if len(keyStr) == 0 {
				return
			}
			fmt.Fprintf(&inputBuilder, "%s %d\n", keyStr, val)
			allItems = append(allItems, kvheap.KeyVal{Key: keyStr, Val: val})
		}

		addKeyVal(key1, val1)
		addKeyVal(key2, val2)
		addKeyVal(key3, val3)

		if len(allItems) == 0 {
			t.Skip() // No valid items
		}

		// Run ToySort
		r := bytes.NewReader(inputBuilder.Bytes())
		result := ToySort(r, n)

		// Property 1: Result length should be min(n, len(allItems))
		expectedLen := min(len(allItems), n)
		if len(result) != expectedLen {
			t.Errorf("expected %d items, got %d", expectedLen, len(result))
		}

		// Property 2: Result should be sorted in descending order by value
		for i := 0; i < len(result)-1; i++ {
			if result[i].Val < result[i+1].Val {
				t.Errorf("result not sorted descending: result[%d].Val=%d < result[%d].Val=%d",
					i, result[i].Val, i+1, result[i+1].Val)
			}
		}

		// Property 3: Result should contain the top N items
		// Sort allItems descending to get expected top N
		sort.Slice(allItems, func(i, j int) bool {
			return allItems[i].Val > allItems[j].Val
		})
		expectedTopN := allItems
		if len(allItems) > n {
			expectedTopN = allItems[:n]
		}

		// Check that result values match expected top N values
		if len(result) != len(expectedTopN) {
			t.Errorf("result length %d != expected length %d", len(result), len(expectedTopN))
		}
		resultVals := make([]int64, len(result))
		for i, kv := range result {
			resultVals[i] = kv.Val
		}
		expectedVals := make([]int64, len(expectedTopN))
		for i, kv := range expectedTopN {
			expectedVals[i] = kv.Val
		}
		for i := range resultVals {
			if i >= len(expectedVals) {
				break
			}
			if resultVals[i] != expectedVals[i] {
				t.Errorf("result[%d].Val=%d != expected[%d].Val=%d", i, resultVals[i], i, expectedVals[i])
			}
		}
	})
}
