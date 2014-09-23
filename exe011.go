package main

import (
  "fmt"
  "math"
)

func main() {
	var i,j,k,count int = 0,0,0,0
	for i = 101; i<=200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == k+1 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("total", count)
}