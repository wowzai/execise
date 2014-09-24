package main

import "fmt"

func main() {
	var m,n,r,x int
	fmt.Printf("请输入两个正整数：")
	fmt.Scanf("%d%d",&m,&n)
	x = m * n
	for n != 0 {
		r = m % n
		m = n
		n = r
	}
	fmt.Printf("最大公约数：%d 最小公倍数：%d\n",m, x/m)
}