package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				if (i != j) && (i != k) && (j != k) {
					fmt.Println(i, j, k)
				}
			}
		}
	}
}
