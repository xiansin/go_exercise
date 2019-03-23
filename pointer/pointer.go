package main

import "fmt"

func main()  {
    var ptr *int
    i := 1
    ptr = &i
    fmt.Println(ptr,*ptr)

    str := "string"
    fmt.Println(str)
    ptr2 := &str
    *ptr2 = "aaa"
    fmt.Println(str)
    /*
    1.指针不能得到一个文字或常量的地址
    2.指针不能进行运算比如 c = *ptr++ 是不被允许的
    3.对一个空指针的反向引用是不合法的，并且会使程序崩溃
    nil->表示 指针，通道，func，接口，映射或切片类型的零值
     */
    //var p *int = nil
    //*p = 0
    // error: invalid memory address or nil pointer dereference
}
