package main

import (
	"fmt"
	"time"
	"timer"
)

type Person struct {
	name string
	age  int
}

func (p *Person) ShowAge(t time.Time, arg ...interface{}) {
	oldAge, ok := arg[0].(int)
	if !ok {
		panic("balabala")
	}
	fmt.Printf("ShowAge time=%v\n", t)
	fmt.Printf("name %s, now age=%d, oldAge=%d\n", p.name, p.age, oldAge)
}

func main() {
	p1 := Person{name: "xiaoming", age: 10}
	tf := timer.NewTimerFunc(time.Second*2, p1.ShowAge, p1.age)
	p1.age = 11

	time.Sleep(time.Second * 3) //let timer work
	if !tf.Stop() {
		fmt.Printf("timer have run, can't stop the timer")
	}
}
