package main

import "fmt"

var var1 = 10
var var2 = 5
var x1, x2, x3 int

func main() {
	x1, x2, x3 = Multiply3Nums(var1, var2)
	Print()
	x1, x2, x3 = Multiply3Nums1(var1, var2)
	Print()
}

func Print() {
	fmt.Printf("x1=%d,x2=%d,x3=%d\n", x1, x2, x3)
}

func Multiply3Nums(var1, var2 int) (int, int, int) {
	return var1 + var2, var1 * var2, var1 - var2
}

func Multiply3Nums1(var1, var2 int) (x1 int, x2 int, x3 int) {
	x1, x2, x3 = var1+var2, var1*var2, var1-var2
	return

}
