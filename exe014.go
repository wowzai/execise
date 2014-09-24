package main

import "fmt"

type B bool

func main() {
	var score int = 0 
	fmt.Printf("请输入分数：")
	fmt.Scanf("%d",&score)
	fmt.Println(B(score >= 90).C("A",B(score >=60).C("B","C")))
}

func (b B) C(t,f interface{}) interface{} {
	if bool(b) == true {
		return t
	}
	return f
}
