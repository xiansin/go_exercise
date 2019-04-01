package main

import "fmt"

func main() {
	//forChart(5)
	//loop()
	//fizzBuzz()
	//forRange()

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d %p", v,*v)
		v = 5
	}
}

func forChart(num int) {
	for i := 0; i <= num; i++ {
		for j := 0; j < i; j++ {
			print("G")
		}
		println()
	}
}

func loop() {
	i := 0
START:
	fmt.Println(i)
	i++
	if i < 15 {
		goto START
	}

}

func fizzBuzz() {
	const (
		FIZZ     = 3
		BUZZ     = 5
		FIZZBUZZ = FIZZ * BUZZ
	)
	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		}
	}
}

func forRange()  {
	str := "Go语言设计的非常棒"
	for index,value := range str{
		fmt.Printf("当前遍历到%d个字符,值为->%c\n",index,value)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
    	fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
    	fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}