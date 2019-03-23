package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main(){
	// 单引号只能赋值一个长度字符 双引号则可以赋值多个
	// 在循环中使用加号 + 拼接字符串并不是最高效的做法，更好的办法是使用函数 strings.Join()
	// 有没有更好地办法了？有！使用字节缓冲（bytes.Buffer）拼接更加给力
	single := 'a'
	str := "jfdhjfsj\n"+"fafafa"

	fmt.Printf("%c\n",single)
	fmt.Print(str)
	rune()
	prefixSuffix()
}

func rune()  {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))
}

func prefixSuffix(){
	str := "This is an example of a string"
	fmt.Printf("这个字符串%s以%s为开头,",str,"Th")
	fmt.Printf("%t\n",strings.HasPrefix(str,"Th"))
	fmt.Printf("这个字符串%s以%s为结尾,",str,"ing")
	fmt.Printf("%t\n",strings.HasSuffix(str,"ing"))

}