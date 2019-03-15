package main

import (
    "fmt"
    "math/rand"
    "time"
)

/**
    1.rand.Float32 和 rand.Float64 介于[0.0,1.0)
    2.rand.Intn(n) 介于 [0,n) 之间的随机数
    3.可以使用 Seed(value) 函数来提供伪随机数的生成种子，一般情况下都会使用当前时间的纳秒级数字
 */
func main() {
    for i := 0; i < 5; i++ {
        r := rand.Int()
        fmt.Printf("%d / ", r)
    }
    fmt.Println()
    for i:=0;i<5 ;i++  {
        r := rand.Intn(8)
        fmt.Printf("%d / ", r)
    }
    timens := int64(time.Now().Nanosecond())
    rand.Seed(timens)
    fmt.Println()
    for i := 0;i<5 ;i++  {
        fmt.Printf("%2.5f / ", 100*rand.Float32())
    }
}
