package main

import (
	"fmt"
	"math/rand"
)

type Vector struct {
	elems []int
	lth   int
	size  int
}

func (v *Vector) expand() {
	sizet := len(v.elems) + 1
	if sizet > v.size {
		v.size = v.size << 1
		newElems := make([]int, v.size)
		copy(newElems, v.elems)
		v.elems = newElems
	}
}

func (v *Vector) Insert(elem, index int) {
	lth := len(v.elems)
	v.expand()
	for i := lth; i > index; i-- {
		v.elems[i] = v.elems[i-1]
	}
	v.elems[index] = elem
	v.lth++
}

func (v *Vector) Remove(lo, hi int) {
	if lo < 0 || hi > len(v.elems) {
		fmt.Printf("Remove err: index out of range\n")
	}
	if lo > hi {
		fmt.Printf("lo:%d litter than hi:%d\n", lo, hi)
	}
	for i := lo; i < v.lth; i++ {
		v.elems[i] = v.elems[i+hi-lo+1]
	}
	v.lth -= hi - lo + 1
}

func (v *Vector) RemoveOne(index int) {
	v.Remove(index, index)
}

func (v *Vector) find(elem, lo, hi int) int {
	if lo < 0 || hi > len(v.elems) {
		fmt.Printf("Remove err: index out of range\n")
	}
	if lo > hi {
		fmt.Printf("lo:%d litter than hi:%d\n", lo, hi)
	}
	for i := lo; i < hi; i++ {
		if v.elems[i] == elem {
			return i
		}
	}
	return -1
}

func (v *Vector) Find(elem int) (index int) {
	index = v.find(elem, 0, v.lth)
	return index
}

func (v *Vector) Deduplicate() {
	oldLth := v.lth
	//必须从后往前，从前往后的话如果删除多个数据，
	//vector的length会变小且会被0覆盖导致后面的元素
	//遍历不到
	for i := oldLth - 1; i > 0; i-- {
		idx := v.find(v.elems[i], 0, i-1)
		fmt.Printf("index:%d elem:%d\n", i, v.elems[i])
		if idx > 0 {
			fmt.Println("idx:", idx)
			v.RemoveOne(idx)
		}
	}
}

type visit func(i *int) //传指针才能改变原先的值

func (v *Vector) Traverse(vt visit) {
	for i := 0; i < v.lth; i++ {
		vt(&v.elems[i])
	}
}

func increase(i *int) {
	(*i)++
}

func double(i *int) {
	(*i) = (*i) << 1
}

func (v *Vector) String() string {
	str := fmt.Sprintf("[")
	for i := 0; i < v.lth; i++ {
		str += fmt.Sprintf("%d ", v.elems[i])
	}
	str += fmt.Sprintf("\b]")
	return str
}

func main() {
	v := new(Vector)
	elems := []int{8, 5, 45, 6, 9, 12}
	v.elems = elems
	v.size = len(elems)
	v.lth = len(elems)
	fmt.Printf("vector:%s\n", v)
	//v.expand()
	v.Insert(18, 3)
	fmt.Printf("vector:%s\n", v)
	v.Remove(2, 5)
	fmt.Printf("vector:%s\n", v)
	v.RemoveOne(2)
	fmt.Printf("vector:%s\n", v)
	for i := 0; i < 5; i++ {
		v.Insert(rand.Intn(100), i)
	}
	fmt.Printf("vector:%s\n", v)
	index := v.Find(5)
	fmt.Printf("index:%d\n", index)
	index = v.Find(15)
	fmt.Printf("index:%d\n", index)
	v.Insert(8, 2)
	v.Insert(5, 4)
	fmt.Printf("vector:%s\n", v)
	v.Deduplicate()
	fmt.Printf("vector:%s\n", v)
	v.Traverse(increase)
	fmt.Printf("vector:%s\n", v)
	v.Traverse(double)
	fmt.Printf("vector:%s\n", v)

}
