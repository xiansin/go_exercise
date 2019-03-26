package main

import (
    "fmt"
    "strconv"
)


func main() {
    ifElse()
}

func ifElse() {
    if true {
        fmt.Println("真")
    } else { // else 是不可进行换行的，换行将会导致编译错误
        fmt.Println("假")
    }
    // 可以初始化一个变量
    if val := 11; val > 10 {
        fmt.Println("大于10")
    }
    checkFuncReturn()
}

func checkFuncReturn()  {
    var orig = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		//os.Exit(1)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

}
