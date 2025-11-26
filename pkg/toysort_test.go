package pkg_test

import (
	"container/heap"
	"fmt"
	"testing"

	. "github.com/masterfuzz/toysort/pkg"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := &KeyValHeap{
		KeyVal{
			Key: "a",
			Val: 10,
		},
		KeyVal{
			Key: "a",
			Val: 6,
		},
	}

	heap.Init(h)
	heap.Push(h, KeyVal{Val: 5})
	heap.Push(h, KeyVal{Val: 12})

	assert.Equal(t, (*h)[0].Val, int64(5))
	assert.Equal(t, (*h)[h.Len()-1].Val, int64(12))
}

func TestTopN(t *testing.T) {
	top := NewKVTopN(5)
	for i := range 20 {
		top.Push(KeyVal{Key: fmt.Sprintf("%d", i), Val: int64(i)})
	}

	assert.Equal(t, 5, top.Items.Len())
	fmt.Printf("%v", top.Items)
	assert.Equal(t, int64(15), top.Pop().Val)

	assert.Equal(t, []string{"19", "18", "17", "16"}, top.TopN())
}
