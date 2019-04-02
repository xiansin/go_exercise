package main

import "fmt"

func main() {
	//forChart(5)
	//loop()
	//fizzBuzz()
	//forRange()
	//bitwiseComplement()

	/*for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d %p\n", v, &v)
		v = 5
	}*/

	/*for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}*/
	/*for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}
	*/
	/*s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}*/
	/*for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}*/
	tag()
}
func tag(){
	LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
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
		if i < 15{
			goto START
		}
}

func bitwiseComplement(){
	for i :=0; i<10;i++  {
		fmt.Printf("the complement of %b is: %b\n",i,^i)
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