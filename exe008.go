package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j) /*-3d 表示左对齐 占三位 */
		}
		fmt.Println()
	}
}
