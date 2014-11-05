package main

import "fmt"

func main() {
	var array [10]int
	/*
		for i := 0; i < 10; i++ {
			fmt.Printf("array[%d]=",i)
			fmt.Scanf("%d\n",&array[i])
		}
		fmt.Println()
	*/
	array = [10]int{8, 9, 14, -1, 3, 27, 33, 45, 19, 23}
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
		}
	}
	fmt.Println(array)
}
