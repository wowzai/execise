package main

import "fmt"

func recJX(x int) int {
	if x < 0 {
		fmt.Printf("您输入的值不符合规范！\n")
		return -1
	}
	if x > 1 {
		return x * recJX(x-1)
	} else {
		return 1
	}
}

func main() {
	jx5 := recJX(5)
	fmt.Printf("5! = %d\n", jx5)
}
