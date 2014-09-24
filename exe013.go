package main

import "fmt"

func main() {
	var n,i int = 0,0
	fmt.Println("Please input a number:")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("%d=",n)
	for i=2; i<=n; i++ {
		for n != i {
			if n % i == 0 {
				fmt.Printf("%d*",i)
				n = n / i
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n",n)
}