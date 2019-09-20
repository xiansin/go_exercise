/*
 1.使用 defer 关键字可以解锁加锁资源
 2.可以在最后关闭打开文件
 3.打印最终报告
 4.关闭数据库连接
 5.最终代码运行情况
 有点类似PHP __destruct() 函数
 */
package main

import (
    "fmt"
    "log"
    "io"
)

func main() {
    fun1()
    //fun3()
    //fun5()
    //recordParamsReturnVar("Go")
}

func fun1() {
    fmt.Printf("In function1 at the top\n")
    defer fun2()
    fun2()
    fmt.Printf("In function1 at the bottom!\n")
}
func fun2() {
    fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

func fun3() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }
}

func fun4(){
    defer untrace(trace("fun4"))
    fmt.Println("in fun4")
}

func fun5()  {
    defer untrace(trace("fun5"))
    fmt.Println("in fun5")
    fun4()
}

func recordParamsReturnVar(s string) (n int, err error){
    defer func() {
        log.Printf("func1(%q) = %d, %v", s, n, err)
    }()
    return 7, io.EOF
}

func trace(s string) string   {
    fmt.Println("entering:", s)
    return s
}
func untrace(s string) {
    fmt.Println("leaving:", s)
}
