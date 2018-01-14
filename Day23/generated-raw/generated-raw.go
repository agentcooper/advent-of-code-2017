
		// auto-generated
		package main

		import "fmt"

		func main() {
			a,b,c,d,e,f,g,h := 1,0,0,0,0,0,0,0
	
b = 65
c = b
if a != 0 { goto L4 }
if 1 != 0 { goto L8 }
L4: b = b * 100
b = b - -100000
c = b
c = c - -17000
L8: f = 1
d = 2
L10: e = 2
L11: g = d
g = g * e
g = g - b
if g != 0 { goto L16 }
f = 0
L16: e = e - -1
g = e
g = g - b
if g != 0 { goto L11 }
d = d - -1
g = d
g = g - b
if g != 0 { goto L10 }
if f != 0 { goto L26 }
h = h - -1
L26: g = b
g = g - c
if g != 0 { goto L30 }
if 1 != 0 { goto END }
L30: b = b - -17
if 1 != 0 { goto L8 }

		END: fmt.Println(h)
	}
