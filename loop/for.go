package main

import "fmt"

func main(){
	//forChart(5)
	//loop()
	bitwiseComplement()
}

func forChart(num int){
	for i:=0 ; i <= num ;i++  {
		for j:=0;j<i;j++ {
			print("G")
		}
		println()
	}
}

func loop(){
	i := 0
	START:
		fmt.Println(i)
		i++
		if i < 15{
			goto START
		}
}

func bitwiseComplement(){
	for i :=0; i<10;i++  {
		fmt.Printf("the complement of %b is: %b\n",i,^i)
	}
}

