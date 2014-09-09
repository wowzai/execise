package main

import "fmt"

/*
阶梯利润问题
*/

func main() {
	var i float32 = 0.0
	var bonus float32 = 0.0
	fmt.Print("输入利润：")
	fmt.Scanf("%f\n",&i)
	switch {
    // fallthrough 需要显示的写出来，要不然golang
	// 在执行满足的第一个条件后面的语句完后直接跳出
	// switch
	case i > 1000000:
		bonus = (i - 1000000) * 0.01
		i = 1000000
		fallthrough
	case i > 600000:
		bonus = (i - 600000) * 0.015
		i = 600000
		fallthrough
	case i > 400000:
		bonus = (i - 400000) * 0.03
		i = 400000
		fallthrough
	case i > 200000:
		bonus = (i - 200000) * 0.05
		i = 200000
		fallthrough
	case i > 100000:
		bonus = (i - 100000) * 0.075
		i = 100000
		fallthrough
	default:
		bonus += i * 0.1
	}
	fmt.Printf("提成总计：%f\n", bonus)
}