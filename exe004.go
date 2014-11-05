package main

import "fmt"

func main() {
	var y, m, d, days int = 0, 0, 0, 0
	fmt.Scanf("%d%d%d\n", &y, &m, &d)
	fmt.Printf("%d年%d月%d日", y, m, d)
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	fmt.Printf("是今年的第%d天!\n", days)
}
