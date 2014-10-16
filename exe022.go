package main

import "fmt"

func main() {
	var i, j, k int16 // i是a的对手，j是b的对手，k是c的对手
	for i = 'x'; i <= 'z'; i++ {
		for j = 'x'; j <= 'z'; j++ {
			if i != j {
				for k = 'x'; k <= 'z'; k++ {
					if i != k && j != k {
						if i != 'x' && k != 'x' && k != 'z' {
							fmt.Printf("order is a--%c b--%c c--%c\n", i, j, k)
						}
					}
				}
			}
		}
	}
}
