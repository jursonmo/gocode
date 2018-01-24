package ringBuffer

import (
	"testing"
)

func TestGetBuf(t *testing.T) {
	size := 4
	rb := CreateRingBuffer(size)
	bs := make([]*PktBuf, size)
	for i := 0; i < size; i++ {
		b := rb.GetBuf()
		if b == nil {
			t.Fatal("shouldn't be nil")
		}
		bs[i] = b
	}
	buf := rb.GetBuf()
	if buf != nil {
		t.Fatal("should be nil")
	}

	if ok := bs[0].Release(); !ok {
		t.Fatal("Release fail")
	}
	buf = rb.GetBuf()
	if buf == nil {
		t.Fatal("shouldn't be nil")
	}
}

func TestRBLen(t *testing.T) {
	size := 4
	rb := CreateRingBuffer(size)

	if rb.Len() != 4 {
		t.Fatal("rb.Len() fail , %d", rb.Len())
	}

	bs := make([]*PktBuf, size)
	for i := 0; i < size; i++ {
		b := rb.GetBuf()
		if b == nil {
			t.Fatal("shouldn't be nil")
		}
		bs[i] = b
	}
	if rb.Len() != 0 {
		t.Fatal("rb.Len() fail , %d", rb.Len())
	}

	for i := 0; i < 3; i++ {
		if ok := bs[i].Release(); !ok {
			t.Fatal("Release fail")
		}
	}

	if rb.Len() != 3 {
		t.Fatal("rb.Len() fail , %d", rb.Len())
	}
}
