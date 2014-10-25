package main

import "fmt"

func main() {
	for i:=3; i<100; i++ {
		var j = 2
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if i == j {
			fmt.Print(i," ")
		}
	}
}