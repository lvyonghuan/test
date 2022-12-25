package main

import "fmt"

// 先被defer的后执行，且defer在函数return后执行。而函数在打印2的时候已经return了。所以打印3的一步实际上是没有被读取的，自然不会打印3。
// 如果把return抽出来放到最后，那结果将会是2 3 1.
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
