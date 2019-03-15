package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	// 可以使用 println 这样子就可以不用包含 fmt 包
	// println("Hello World!")
	/* 格式化输出
	   %d      格式化整数
	   %0d     输出定长的整数，其中开头数字 0 是必须的
	   %x      和 %X 用于格式化 16 进制表示的数字
	   %g      用于格式化浮点型
	   %f      输出浮点数
	   %e      输出科学计数法
	   %v      表示复数 如果只需要表示复数的一部分可以使用 %f
	   %n.mg   用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 和 f,eg：%5.2e
	*/
	fmt.Printf("%d\n", 1)
	fmt.Printf("%5.4e", 3.4)
}
