package main

import (
	"fmt"
	"math"
)

func match(x int) bool {
	if x/10000 <= 0 {
		fmt.Printf("您输入的五位数不正确\n")
		return false
	}
	var result = true
	for i := 1; i <= 2; i++ {
		m := x % int(math.Pow10(i)) / int(math.Pow10(i-1))
		n := x % int(math.Pow10(5-i+1)) / int(math.Pow10(5-i))
		if m != n {
			result = false
		}
	}
	return result
}

func main() {
	var x int
	fmt.Printf("请输入一个5位数：")
	fmt.Scanf("%d", &x)
	match := match(x)
	if match {
		fmt.Printf("%d是一个回文数\n", x)
	} else {
		fmt.Printf("%d不是一个回文数\n", x)
	}
}
