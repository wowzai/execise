package main

import "fmt"

func main() {
	var i, j, sum = 2.0, 1.0, 0.0
	//var t = 0.0  //中间变量
	for k := 0; k < 20; k++ {
		sum += i / j
		/*
		   t = j  // t当作中间变量缓存j的值
		   j = i
		   i += t
		*/
		// golang中也可以换成如下形式而无需中间变量
		i, j = i+j, i
	}
	fmt.Printf("sum is %f\n", sum)
}
