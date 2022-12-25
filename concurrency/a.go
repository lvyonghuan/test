package main

import "fmt"

// 这里好像没给出具体要做啥。我就理解为用通道让“出现”正常输出了。
func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("出现")
	}()
	<-ch
}
