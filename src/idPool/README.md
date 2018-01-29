### idPool

so simple id pool,  NewIDPool(10), you will have a idpool that id from 1 to 10

```go
  poolSize := 10
  idpool := NewIDPool(poolSize) // id from 1 to 10
  id := idpool.GetID()
  if id == 0 {
    fmt.Println("id pool exhaust")
    return
  }
  // put id back to pool when you didn't need the id
  idpool.PutID(id)
