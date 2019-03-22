package main

import (
    "fmt"
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
}

func checkFuncReturn()  {
    
}
