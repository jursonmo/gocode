package ringBuffer

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type PktBuf struct {
	sync.Mutex
	rb  *RingBuffer
	id  int
	ref int32
	len uint16
	buf [1024]byte
}

type RingBufferHeader struct {
	r, w, count uint16
	//size uint32
}

type RingBuffer struct {
	getLock sync.Mutex
	putLock sync.Mutex
	hdr     RingBufferHeader
	buf     []*PktBuf //存储指针
}

const (
	BufferNum = 128
)

//var DefaultRB *RingBuffer
//var RingBufferNum = BufferNum + 1

func init() {
	//DefaultRB = CreateRingBuffer(BufferNum)
}

func CreateRingBuffer(n int) *RingBuffer {
	mp := make([]PktBuf, n)
	rb := NewRingBuffer(n + 1)
	for i := 0; i < n; i++ {
		mp[i].id = i
		if RBPutBuf(rb, &mp[i]) == false {
			fmt.Printf("i =%d, RBPutBuf error\n", i)
			panic("RBPutBuf error")
		}
	}
	return rb
}
func NewRingBuffer(size int) *RingBuffer {
	rb := &RingBuffer{
		buf: make([]*PktBuf, size),
	}
	rb.hdr.count = uint16(size)
	return rb
}

func RBPutBuf(rb *RingBuffer, pb *PktBuf) bool {
	rb.putLock.Lock()
	pb.rb = rb
	idx := (rb.hdr.w + 1) % rb.hdr.count
	//fmt.Printf("r=%d,w=%d,idx=%d\n", rb.hdr.r, rb.hdr.w, idx)
	if idx != rb.hdr.r {
		rb.buf[rb.hdr.w] = pb
		rb.hdr.w = idx
		rb.putLock.Unlock()
		return true
	}
	rb.putLock.Unlock()
	return false
}

func RBGetBuf(rb *RingBuffer) *PktBuf {
	rb.getLock.Lock()
	idx := rb.hdr.r
	if idx != rb.hdr.w {
		rb.hdr.r = (rb.hdr.r + 1) % rb.hdr.count
		rb.getLock.Unlock()
		return rb.buf[idx]
	}
	rb.getLock.Unlock()
	return nil
}

func (rb *RingBuffer) GetBuf() *PktBuf {
	pb := RBGetBuf(rb)
	if pb == nil {
		return nil
	}
	pb.HoldPktBuf()
	return pb
}

func (pb *PktBuf) Release() bool {
	if pb == nil {
		panic("pb is nil")
	}
	if pb.rb == nil {
		panic("pb.rb is nil")
	}
	ref := releasePktBuf(pb)
	if ref == 0 {
		return RBPutBuf(pb.rb, pb)
	}
	if ref < 0 {
		panic("ref < 0")
	}
	return false
}

func (pb *PktBuf) HoldPktBuf() {
	atomic.AddInt32(&pb.ref, 1)
}

func releasePktBuf(pb *PktBuf) int32 {
	return atomic.AddInt32(&pb.ref, -1)
}

func (rb *RingBuffer) Len() uint16 {
	//fmt.Println(rb.hdr.w, rb.hdr.r)
	if rb.hdr.w > rb.hdr.r {
		return rb.hdr.w - rb.hdr.r
	}
	return rb.hdr.count - rb.hdr.r + rb.hdr.w
}

func (rb *RingBuffer) Show() {
	fmt.Println("===========show rb==========")
	for i, pb := range rb.buf {
		if pb != nil {
			fmt.Printf("i=%d, pb.id=%d , pb.ref=%d, pb.len=%d\n", i, pb.id, pb.ref, pb.len)
		}
	}
}
