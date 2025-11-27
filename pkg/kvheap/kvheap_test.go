package kvheap_test

import (
	"container/heap"
	"fmt"
	"testing"

	. "github.com/masterfuzz/toysort/pkg/kvheap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := &KeyValHeap{
		KeyVal{
			Key: []byte("a"),
			Val: 10,
		},
		KeyVal{
			Key: []byte("a"),
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
	tokv := func(i int) KeyVal {
		return KeyVal{Key: fmt.Appendf(nil, "%d", i), Val: int64(i)}
	}

	top := NewKVTopN(5)
	for i := range 20 {
		top.Push(tokv(i))
	}

	assert.Equal(t, 5, top.Items.Len())
	fmt.Printf("%v", top.Items)
	assert.Equal(t, int64(15), top.Pop().Val)

	assert.Equal(t, []KeyVal{tokv(19), tokv(18), tokv(17), tokv(16)}, top.TopN())
}
