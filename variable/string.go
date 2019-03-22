package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main(){
    //prefixSuffix()
    //contains()
    //index()
    //calc()
    conv()
}

func prefixSuffix(){
    var str = "This is string"
    fmt.Printf("这个字符串%s是不是以%s开始",str,"Th")
    fmt.Printf("%t\n",strings.HasPrefix(str,"Th"))
    fmt.Printf("这个字符串%s是不是以%s结尾",str,"ing")
    fmt.Printf("%t\n",strings.HasSuffix(str,"ing"))
}

func contains()  {
    str := "string"
    fmt.Printf("%t\n",strings.Contains(str,"s"))
}

func index()  {
    str := "string string"
    fmt.Printf("%d\n",strings.Index(str,"g"))
    fmt.Printf("%d\n",strings.LastIndex(str,"s"))
    fmt.Printf("%d\n",strings.IndexRune(str,99))
}

func calc()  {
    str := " string String "
    fmt.Printf("%d\n",strings.Count(str,"s"))
    fmt.Printf("%s\n",strings.Replace(str,"s","b",1))
    fmt.Printf("%s\n",strings.Repeat(str,2))
    fmt.Printf("%s\n",strings.ToLower(str))
    fmt.Printf("%s\n",strings.ToUpper(str))
    fmt.Printf("%s\n",strings.TrimSpace(str))
    fmt.Printf("%s\n",strings.Trim(str," "))
    fmt.Printf("%s\n",strings.TrimLeft(str," s"))
    fmt.Printf("%s\n",strings.TrimRight(str," "))
    fmt.Printf("%s\n",strings.Fields(str))
    fmt.Printf("%s\n",strings.Split(str,"S"))
    fmt.Printf("%s\n",strings.Join(strings.Split(str,"S"),""))
    fmt.Printf("%p\n",strings.NewReader(str))
}

func conv()  {
    fmt.Printf("%s\n",strconv.Itoa(1))
    fmt.Printf("%d\n",strconv.IntSize)
    an,_ := strconv.Atoi("2")
    fmt.Printf("%d\n",an)
}
