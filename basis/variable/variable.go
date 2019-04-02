package main

import "fmt"

var global = 1

func main(){
    global := 2
    test()
    fmt.Println(global)
    print1()
    VarType()
    Convert()
    Complex()
    String()
}

func test(){
    fmt.Println(global)
    global = 3
}

func print1()  {
    fmt.Println(global)
}
/*
 类型运算
 */
func VarType()  {
    var a int
    var b int32
    a = 15
    // 编译错误，不是同类型
    //b = a + a
    // 编译成功 5 是常量
    b = b + 5

    fmt.Println(a,b)
}
/*
 转换类型
 */
func Convert(){
    var a int16 = 32
    var b int32
    b = int32(a)
    fmt.Printf("16 bit int is %d\n",a)
    fmt.Printf("32 bit int is %d\n",b)
}
/**
 复数：re+imI 来表示，其中 re 代表实数部分，im 代表虚数部分，I 代表根号负 1。
    Complex32
    Complex64
 */
func Complex(){
    var c1 complex64 = 5 + 10i
    fmt.Printf("The value is%v\n",c1)
    fmt.Printf("c1实数%v\n",real(c1))
    fmt.Printf("c1虚数%v\n",imag(c1))
    // cmath 包 里面的相关函数是常用 complex128 作为参数的，在对内存要求不是那么高的时候最好采用 complex128 作为计算类型
    // 如果 re 和 im 都是 float32 类型的 可以使用 c = complex(re, im) 来获得
}

func String()  {
    /*
    判断是否为字母：unicode.IsLetter(ch)
    判断是否为数字：unicode.IsDigit(ch)
    判断是否为空白符号：unicode.IsSpace(ch)
     */
    var ch int = '\u0041'
    var ch2 int = '\u03B2'
    var ch3 int = '\U00101234'
    fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
    fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
    fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
    fmt.Printf("%U - %U - %U", ch, ch2, ch3) // UTF-8 code point
}



