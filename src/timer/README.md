###  refer to https://github.com/siddontang/go/tree/master/time2
##### go 标准库timer.AfterFunc(d, callback) 最终会创建新的goroutine 来执行回调函数callback，如果timer.AfterFun调用过多的话，就会产生很多goroutine,性能下降，时间轮不会对每个timer起新的goroutine 来执行回调函数，这就是性能的关键， 同时runtime timer 用最小堆来做定时器，不如时间轮。但这里实现的时间轮精度比较低的，属于精度低效率高的定时器, 也就是以精度换取效率。精度是可以调节的。 这个完全参照内核的定时器，go来重写。 

#### 下面这个图，是当年学习内核定时器时的草稿。
![image](https://github.com/jursonmo/gocode/raw/master/src/timer/timer.jpg)
