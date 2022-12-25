// 好久没写这种大的模拟类型的东西了，有点麻。没过样例，也有点乱。
// 大概思路是做一个切片，然后每次输入完之后去检查切片里的数据。然后输出不同的。
// 不过第一次输入之后就中断了。写完（？）业务再来debug。（所以大概是不会de了）
package main

import "fmt"

type Num struct {
	quality int
	time    int64
	ship    int
	If      bool
}

var num []Num

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		var temp Num
		var t int64
		var k int
		_, err := fmt.Scanln(&t, &k)
		if err != nil {
			return
		}
		temp.time = t
		temp.ship = i
		temp.If = true
		var Qulity []int
		for j := 0; j < k; j++ {
			var tempQulity int
			_, err := fmt.Scanln(&tempQulity)
			if err != nil {
				return
			}
			var flag = true
			Qulity = append(Qulity, tempQulity)
			for _, value := range Qulity {
				if value == tempQulity {
					flag = false
				}
			}
			if !flag {
				continue
			}
			temp.quality = tempQulity
			flag = Append(temp)
			if flag {
				Qulity = append(Qulity, tempQulity)
			}
		}
		judge(Qulity)
	}
}

func Append(temp Num) bool {
	for _, value := range num {
		if value.If && temp.time-value.time > 86400 {
			value.If = false
		}
		if value.If && value.quality == temp.quality {
			return false
		}
	}
	num = append(num, temp)
	return true
}

func judge(qulity []int) {
	var count = 0
	for _, value := range qulity {
		for _, tempNum := range num {
			if tempNum.If && tempNum.quality != value {
				count++
			}
		}
	}
	fmt.Println(count)
}
