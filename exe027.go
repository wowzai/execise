package main 

import (
	"fmt"
	"os"
	"strconv"
)

func putchar(n int) {
	var c byte
	if n >= 1 {
		fmt.Scanf("%c", &c)
		putchar(n - 1)
		fmt.Printf("%c", c)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("exe027.exe n")
		return
	}
	n,err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	putchar(n)
}
