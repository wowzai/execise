package main

import "fmt"

func main() {
	var i, j, k, m = 0, 0, 0, 100
	for m = 100; m < 1000; m++ {
		i = m / 100
		j = m % 100 / 10
		k = m % 10
		if m == i*i*i+j*j*j+k*k*k {
			fmt.Printf("m=%d\t", m)
		}
	}
	fmt.Println()
}
