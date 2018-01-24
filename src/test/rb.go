package main

import (
	"fmt"
	"ringBuffer"
)

var DataBufferNum int = 4

func main() {
	rb := ringBuffer.CreateRingBuffer(DataBufferNum)
	fmt.Println(rb)

	fmt.Printf("rb.buf len=%d\n", rb.Len())
	rb.Show()

	for j := 0; j < DataBufferNum; j++ {
		pb := rb.GetBuf()
		if pb == nil {
			fmt.Println("pb == nil")
			return
		}
		for i := 0; i < DataBufferNum+j; i++ {
			pb.HoldPktBuf()
			handleBuf(pb)
		}
		pb.Release()
	}
	rb.Show()
	for j := 0; j < DataBufferNum; j++ {
		pb := rb.GetBuf()
		if pb == nil {
			fmt.Println("pb == nil")
			return
		}

		for i := 0; i < DataBufferNum+j; i++ {
			pb.HoldPktBuf()
			handleBuf(pb)
		}
		pb.Release()
	}
	rb.Show()
}

func handleBuf(pb *ringBuffer.PktBuf) {
	pb.Lock()
	//fmt.Printf("handleBuf\n")
	pb.Unlock()

	pb.Release()
}
