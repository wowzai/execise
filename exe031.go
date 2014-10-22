package main

import (
	"fmt"
	"strings"
)

func main() {
	inweek := []string{"Monday","Tuesday","Wednesday","Thurday","Friday","Saturday","Sunday"}
	instr := ""
	fmt.Printf("input a char:")
	fmt.Scanf("%d\n", &instr)
	outweek(instr, inweek)
}

func outweek(instr string, inweek []string) {
	nextstr := ""
	nextweek := make([]string,0)
	for _,value := range inweek {
		index := strings.Index(value, instr)
		if index == 0 {
			nextweek = append(nextweek, value)
		}
	}
	fmt.Printf("%s\n", nextweek[0:])
	if len(nextweek) > 1 {
		fmt.Printf("input next char:")
		fmt.Scanf("%s\n", &nextstr)
		nextstr = instr + nextstr
		outweek(nextstr, nextweek)
	}
}