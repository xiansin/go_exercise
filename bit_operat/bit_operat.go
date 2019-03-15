package main

import "fmt"

type ByteSize float64

const (
    _ = iota // 空白符用于计数
    KB ByteSize = 1<<(10*iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main()  {
    BitOpera()
}

func BitOpera()  {
    // 二元运算
    fmt.Printf("1 & 1=%d\n",1&1)
    fmt.Printf("1 | 1=%d\n",1|1)
    fmt.Printf("1 ^ 1=%d\n",1^1)
    // 一元运算
    fmt.Printf("^1=%d\n",^1)
    fmt.Printf("1<<1=%d\n",1<<1)
    fmt.Printf("1>>1=%d\n",1>>1)
}