# gmmpool

A multi level memory pool for Golang:
learn from github.com/liudanking/gmmpool

```go
package main

import (
	"bytes"
	"log"

	"gmmpool"
)

func main() {
	pool := gmmpool.NewMultiLevelPool([]gmmpool.PoolOpt{
		gmmpool.PoolOpt{Num: 10, Size: 1024},     // level 0
		gmmpool.PoolOpt{Num: 10, Size: 1024 * 2}, // level 1
	})

	buf := pool.Get(1025)
	b := []byte{1, 2, 4, 5, 6, 6}
	data, err := buf.ReadAll(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(data)

	pool.Put(buf)
	log.Print(pool) //log.Print(pool.String())
	log.Println("pool leak:", pool.IsLeak())
}


```



