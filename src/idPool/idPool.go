package idPool

import (
	"fmt"
	"sync"
)

type IDPools struct {
	sync.Mutex
	idPool []int
}

var DefaultIDPool *IDPools

func init() {
	DefaultIDPool = NewIDPool(1024) //1 - 1024
}

func NewIDPool(size int) *IDPools {
	idpool := &IDPools{
		idPool: make([]int, size+1), //id from 1 to  szie
	}
	idpool.initIDPool()
	return idpool
}

func (idp *IDPools) initIDPool() {
	for id := 1; id < len(idp.idPool); id++ {
		idp.PutID(id)
	}
}

func (idp *IDPools) PutID(id int) {
	idp.Lock()
	idp.idPool[id] = idp.idPool[0]
	idp.idPool[0] = id
	idp.Unlock()
}

func (idp *IDPools) GetID() int {
	idp.Lock()
	id := idp.idPool[0]
	idp.idPool[0] = idp.idPool[id]
	idp.Unlock()
	return id
}

func GetID() int {
	return DefaultIDPool.GetID()
}

func PutID(id int) {
	DefaultIDPool.PutID(id)
}

func (idp *IDPools) ShowIDPool() {
	fmt.Println("show id pool:", idp.idPool)
}
