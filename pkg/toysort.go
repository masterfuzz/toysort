package pkg

import "container/heap"

type KeyVal struct {
	Key string
	Val int64
}

type KeyValHeap []KeyVal

var _ heap.Interface = &KeyValHeap{}

func (h KeyValHeap) Len() int {
	return len(h)
}

func (h KeyValHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h KeyValHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *KeyValHeap) Push(x any) {
	*h = append(*h, x.(KeyVal))
}

func (h *KeyValHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KVHeapTopN struct {
	Items *KeyValHeap
	N int
}

func NewKVTopN(n int) KVHeapTopN {
	h := &KeyValHeap{}
	heap.Init(h)
	return KVHeapTopN{
		Items: h,
		N: n,
	}
}

func (h KVHeapTopN) Push(x KeyVal) {
	heap.Push(h.Items, x)
	if h.Items.Len() > h.N {
		heap.Pop(h.Items)
	}
}

func (h KVHeapTopN) TopN() []string {
	res := make([]string, h.Items.Len())
	i := h.Items.Len() - 1
	for i >= 0 {
		res[i] = h.Pop().Key
		i--
	}
	return res
}

func (h KVHeapTopN) Pop() KeyVal {
	return heap.Pop(h.Items).(KeyVal)
}
