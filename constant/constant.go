package main

const a = 1
const b,c  = 2,3
const (
    i = iota
    j
    k
)

func main(){
    // 不能对常量进行重新赋值
    //j = 1
    println(j)
}