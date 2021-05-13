package utils

import (
	"fmt"
	"math"
)

// MinHeap contains the struct of Url
type MinHeap struct {
	Element []*Url
}

// Url contains the frequence of the url and the url itself.
type Url struct {
	Freq int64
	Addr string
}

// NewMinHeap constructs a new Minheap which has a url with minimum frequence and it returns the address of the Minheap.
func NewMinHeap() *MinHeap {
	head := &Url{math.MinInt64, "None"}
	h := &MinHeap{Element: []*Url{head}}
	return h
}

// Length returns the real length of Minheap without the init one.
func (H *MinHeap) Length() int {
	return len(H.Element) - 1
}

// Min gets the minimum of the Minheap.
func (H *MinHeap) Min() (*Url, error) {
	if len(H.Element) > 1 {
		return H.Element[1], nil
	}
	return nil, fmt.Errorf("heap is empty")
}

// Insert inserts items requires ensuring the nature of the Minheap.
func (H *MinHeap) Insert(v *Url) {
	H.Element = append(H.Element, v) //append the element to the Minheap.
	i := len(H.Element) - 1
	for ; (H.Element[i/2]).Freq > v.Freq; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// DeleteMin deletes and return the minimum,and the element left is still a Minheap.
func (H *MinHeap) DeleteMin() (*Url, error) {
	if len(H.Element) <= 1 {
		return nil, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1].Freq < H.Element[child].Freq {
			child++
		}
		if lastElement.Freq > H.Element[child].Freq {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}
