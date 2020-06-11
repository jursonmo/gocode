package main

import (
	"fmt"
	"sort"
)

type IntTop []int

func (p *IntTop) Len() int {
	return len(*p)
}
func (p *IntTop) Less(i, j int) bool {
	return (*p)[i] < (*p)[j] //最小堆
}
func (p *IntTop) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *IntTop) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *IntTop) Pop() (v interface{}) {
	*p, v = (*p)[:p.Len()-1], (*p)[p.Len()-1]
	return
}

func (p IntTop) CanPush(x interface{}) bool {
	n := x.(int)
	if n > p[0] { //最小堆，如果大于p[0]，说明属于topk, 如果求lowK, 这里要判断n < p[0]
		p[0] = n //x replace p[0], and then down()
		return true
	}
	return false
}

func (p IntTop) TopCap() int {
	return 4 //i wanna top4
}

func main() {
	ih := &IntTop{}
	err := Init(ih)
	if err != nil {
		fmt.Println(err)
		return
	}
	Push(ih, 8)
	fmt.Println("after push 8", ih)
	Push(ih, 1)
	fmt.Println("after push 1", ih)
	Push(ih, 7)
	fmt.Println("after push 7", ih)
	Push(ih, 4)
	fmt.Println("after push 4", ih)
	Push(ih, 6)
	fmt.Println("after push 6", ih)
	Push(ih, 5)
	fmt.Println("after push 5", ih)
	Push(ih, 6)
	fmt.Println("after push 6", ih)
	Push(ih, 13)
	fmt.Println("after push 13", ih)
}

type Interface interface {
	sort.Interface
	Push(x interface{})         // add x as element Len()
	Pop() interface{}           // remove and return element Len() - 1.
	TopCap() int                //if you want top10, return 10
	CanPush(x interface{}) bool //can Push , compare with heapTop[0]
}

func Init(h Interface) error {
	// heapify
	n := h.Len()
	if n > h.TopCap() {
		return fmt.Errorf("h.Len():%d > h.TopCap():%d", h.Len(), h.TopCap())
	}
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
	return nil
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	if h.Len() < h.TopCap() {
		h.Push(x)
		up(h, h.Len()-1)
		return
	}

	if h.Len() != h.TopCap() {
		panic("h.Len() != h.TopCap()")
	}

	if h.CanPush(x) {
		down(h, 0, h.Len())
	}
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
