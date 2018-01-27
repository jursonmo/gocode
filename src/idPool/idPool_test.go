package idPool

import "testing"

func TestDefaultGetID(t *testing.T) {
	id := GetID()
	if id != 1024 {
		t.Fatal("id != 1024, is %d ", id)
	}

	PutID(id)
	id = GetID()
	if id != 1024 {
		t.Fatal("id != 1024, is %d ", id)
	}
}

func TestGetID(t *testing.T) {
	poolSize := 10
	idpool := NewIDPool(poolSize)
	id := idpool.GetID()
	if id != 10 {
		t.Fatal("id  is %d ", id)
	}

	idpool.PutID(id)

	ids := make([]int, poolSize)
	for i := 0; i < poolSize; i++ {
		ids[i] = idpool.GetID()
		if ids[i] == 0 {
			t.Fatal("id  is %d ", ids[i])
		}
	}
	idpool.ShowIDPool()

	id = idpool.GetID()
	if id != 0 {
		t.Fatal("id  is %d ", id)
	}

}
