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

func (v *Vector) bubbleSort(lo, hi int) bool {
	var sorted = true
	for i := lo; i < hi; i++ {
		if v.elems[i] > v.elems[i+1] {
			sorted = false
			v.elems[i], v.elems[i+1] = v.elems[i+1], v.elems[i]
			/*
				t := elems[i]
				elems[i] = elems[i+1]
				elems[i+1] = t
			*/
		}
	}
	return sorted
}

func (v *Vector) bubbleSort2(lo, hi int) int {
	var last = lo
	for i := lo; i < hi; i++ {
		if v.elems[i] > v.elems[i+1] {
			last = i + 1
			v.elems[i], v.elems[i+1] = v.elems[i+1], v.elems[i]
		}
	}
	return last
}

func (v *Vector) Sort() {
	hi := v.lth - 1
	for !v.bubbleSort(0, hi) {
		hi--
	}
}

func (v *Vector) Sort2() {
	hi := v.lth - 1
	for hi > 0 {
		hi = v.bubbleSort2(0, hi)
		fmt.Printf("hi:%d\n", hi)
	}
}

//O(n) 只要遍历一次即可
func (v *Vector) Uniquify() {
	//v.Sort()
	var lo, hi int
	for lo, hi = 0, 0; hi < v.lth; hi++ {
		if v.elems[lo] != v.elems[hi] {
			lo++
			v.elems[lo] = v.elems[hi]
		}
	}
	lo++
	v.lth = lo
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

func BinSearch(elems []int, elem, lo, hi int) int {
	var mi int
	for lo < hi {
		mi = (lo + hi) >> 1
		if elem < elems[mi] {
			hi = mi
		} else if elem > elems[mi] {
			lo = mi + 1
		} else {
			return mi
		}
		/*
			if elem == elems[mi] {
				return mi
			} else if elem < elems[mi] {
				hi = mi
				//recursion
				//return BinSearch(elems, elem, lo, mi)
			} else { //elem > elems[mi]
				lo = mi + 1
				//recursion
				//return BinSearch(elems, elem, mi+1, hi)
			}
		*/
	}
	return -1
}

type Fib struct {
	fib []int
	idx int
	lth int
}

func (f *Fib) get() int {
	return f.fib[f.idx]
}

func (f *Fib) prev() {
	f.idx--
}

func fib(lth int) (fib Fib) {
	if lth < 1 {
		fmt.Printf("lth:%d is err\n", lth)
	}
	elems := make([]int, lth)
	if lth > 2 {
		elems[0] = 1
		elems[1] = 1
		for i := 2; i < lth; i++ {
			elems[i] = elems[i-1] + elems[i-2]
		}
	} else {
		for i := 0; i < lth; i++ {
			elems[i] = 1
		}
	}
	fib.fib = elems
	fib.idx = lth - 1
	fib.lth = lth
	return fib
}

func BinSearchB(elems []int, elem, lo, hi int) int {
	var mi int
	for hi-lo > 1 {
		mi = (lo + hi) >> 1
		if elem < elems[mi] {
			hi = mi
		} else {
			lo = mi
		}
	}
	if elems[lo] == elem {
		return lo
	} else {
		return -1
	}
}

func BinSearchC(elems []int, elem, lo, hi int) int {
	var mi int
	for lo < hi {
		mi = (lo + hi) >> 1
		if elem < elems[mi] {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo - 1
}

func FibSearch(elems []int, elem, lo, hi int) int {
	fib := fib(hi - lo)
	var mi int
	for lo < hi {
		for hi-lo < fib.get() {
			fib.prev()
		}
		mi = lo + fib.get() - 1 //黄金比例切分
		if elem < elems[mi] {
			hi = mi
		} else if elem > elems[mi] {
			lo = mi + 1
		} else {
			return mi
		}
	}
	return -1
}

/*
大规模：插值查找
中规模：折半查找
小规模：顺序查找
*/

func InsertSearch(elems []int, elem, lo, hi int) int {
	var mi int
	for lo < hi {
		mi = lo + (hi-lo)*(elem-elems[lo])/(elems[hi]-elems[lo])
		if elem < elems[mi] {
			hi = mi
		} else if elem > elems[mi] {
			lo = mi + 1
		} else {
			return mi
		}
	}
	return -1
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
	elems = []int{3, 3, 3, 5, 5, 5, 5, 8, 8, 8, 8, 8, 13, 13, 13, 13, 13, 13}
	v.elems = elems
	v.lth = len(elems)
	fmt.Printf("vector:%s\n", v)
	v.Uniquify()
	fmt.Printf("vector:%s\n", v)
	elems = []int{2, 4, 5, 7, 8, 9, 12, 15}
	fmt.Printf("vector:%v\n", elems)
	index = BinSearch(elems, 8, 0, 7)
	fmt.Printf("BinSearch index:%d\n", index)
	index = BinSearch(elems, 1, 0, 7)
	fmt.Printf("BinSearch index:%d\n", index)
	index = FibSearch(elems, 8, 0, 7)
	fmt.Printf("FibSearch index:%d\n", index)
	index = FibSearch(elems, 1, 0, 7)
	fmt.Printf("FibSearch index:%d\n", index)
	index = BinSearchB(elems, 8, 0, 7)
	fmt.Printf("BinSearchB index:%d\n", index)
	index = BinSearchB(elems, 1, 0, 7)
	fmt.Printf("BinSearchB index:%d\n", index)
	index = BinSearchC(elems, 8, 0, 7)
	fmt.Printf("BinSearchC index:%d\n", index)
	index = BinSearchC(elems, 1, 0, 7)
	fmt.Printf("BinSearchC index:%d\n", index)
	index = InsertSearch(elems, 8, 0, 7)
	fmt.Printf("InsertSearch index:%d\n", index)
	index = InsertSearch(elems, 1, 0, 7)
	fmt.Printf("InsertSearch index:%d\n", index)
	elems = []int{22, 8, 7, 10, 19, 12, 3, 2, 11}
	v.elems = elems
	v.lth = len(elems)
	fmt.Printf("before sort vector:%v\n", elems)
	v.Sort()
	fmt.Printf("after sort vector:%v\n", elems)
	elems = []int{1, 2, 3, 4, 8, 7, 10, 19, 12, 11}
	v.elems = elems
	v.lth = len(elems)
	fmt.Printf("before sort2 vector:%v\n", elems)
	v.Sort2()
	fmt.Printf("after sort2 vector:%v\n", elems)

}
