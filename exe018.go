package main

import "fmt"

func main() {
	var temp []int = make([]int, 100)
	var k, p int = 0, 0
	for n := 2; n < 1000; n++ {
		m := n
		for i := 1; i < n; i++ {
			if n%i == 0 {
				if temp[k] == -1 {
					p = k
				}
				m -= i
				temp[k] = i
				k++
			}
		}
		if m == 0 {
			temp[k] = -1
			k++
			fmt.Printf("%v --- %d\n", temp[p:k-1], n)
		} else {
			k = p
		}
	}
}
