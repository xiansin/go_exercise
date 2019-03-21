package main

import (
    "fmt"
    "strings"
)

func main(){
    prefixSuffix()
}

func prefixSuffix(){
    var str string = "This is string"
    fmt.Printf("这个字符串%s是不是以%s开始",str,"Th")
    fmt.Printf("%t\n",strings.HasPrefix(str,"Th"))
    fmt.Printf("这个字符串%s是不是以%s结尾",str,"ing")
    fmt.Printf("%t\n",strings.HasSuffix(str,"ing"))
}
