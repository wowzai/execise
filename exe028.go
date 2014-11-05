package main

import "fmt"

func age(n int) int {
	if n < 1 {
		fmt.Printf("您输入的参数不符合规范\n")
		return -1
	}
	if n == 1 {
		return 10
	} else {
		return 2 + age(n-1)
	}
}

func main() {
	age := age(5)
	fmt.Printf("第五个人的年龄为%d\n", age)
}
