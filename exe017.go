package main

import "fmt"

func main() {
	var a,n,count int
	var sn,tn int
	fmt.Printf("please input a and n")
	fmt.Scanf("%d%d\n", &a, &n)

	for count < n {
		tn += a
		sn += tn
		a *= 10
		count++
	}

	fmt.Printf("a+aa+aaa+...=%d\n",sn)
}