package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Printf("请输入一个数字：")
	fmt.Scanf("%d\n", &x)
	for i := 5; i > 0; i-- {
		r := x / int(math.Pow10(i-1))
		if r > 0 {
			fmt.Printf("%d是一个%d位数字。\n", x, i)
			out(i, x)
			fmt.Printf("\n")
			break
		}
	}
}

func out(n, x int) {
	if n > 1 {
		out(n-1, x)
	}
	// x % (10^n) / (10^n-1)
	r := x % int(math.Pow10(n)) / int(math.Pow10(n-1))
	fmt.Printf("%d", r)
}
