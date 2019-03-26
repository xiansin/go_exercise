package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var1 := 11
	/*
	switch 切换到某一个分支之后不在往下执行，也就是说不需要break
	如果需要继续往下执行 需要使用
	fallthrough
	 */

	switch var1 {
	case 1 * 1:
		fmt.Println("var1 value is 1")
		// fallthrough
	case 2:
		fmt.Println("var1 value not 2")
	case 3, 4:
		fmt.Println("var1 value is 3 or 4")
	default:
		fmt.Println("var1 value is NaN")
	}
	var2 := 5 * rand.Intn(2)
	fmt.Println(var2)
	switch {
	case var2 < 5:
		fmt.Println("var1 value less 5")
	case var2 >= 5:
		fmt.Println("var1 value beyond 5")
	default:
		fmt.Println("var1 value is NaN")
	}
}
