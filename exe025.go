package main

import "fmt"

func main() {
	s,t := 0,1
	for n := 1; n <= 20; n++ {
		t *= n
		s += t
	}
	fmt.Printf("1! + 2! + 3! + ... + 20! = %d\n",s)
}