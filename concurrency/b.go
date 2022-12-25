package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 1)
var wg sync.WaitGroup

func main() {
	wg.Add(3)
	ch <- 1
	go Work("goroutine1")
	ch <- 1
	go Work("goroutine2")
	ch <- 1
	go Work("goroutine3")
	wg.Wait()
	fmt.Println("successful")
}
func Work(workName string) {
	time.Sleep(time.Second)
	<-ch
	// 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
}
