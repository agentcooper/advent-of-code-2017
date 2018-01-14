package main

import "fmt"

func main() {
	b, c, h := 106500, 106500+17000, 0

	for {
		f := 1
		d := 2

	outer:
		for d != b {
			e := 2
			for e != b {
				if d*e == b {
					f = 0
					break outer // manual optimization
				}
				e++
			}
			d++
		}
		if f == 0 {
			h++
			fmt.Printf("h = %d, b = %d, c = %d\n", h, b, c)
		}
		if b == c {
			break
		} else {
			b = b + 17
		}
	}

	fmt.Println(h)
}
