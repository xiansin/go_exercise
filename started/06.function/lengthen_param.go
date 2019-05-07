package main

import "fmt"

func main(){
    x := min()
    fmt.Printf("The minimum is: %d\n", x)
    x = min(5,2,32,1,0)
    fmt.Printf("The minimum is: %d\n", x)
    x = min(1, 3, 2, 0)
    fmt.Printf("The minimum is: %d\n", x)
    slice := []int{7,9,3,5,1}
    x = min(slice...)
    fmt.Printf("The minimum in the slice is: %d", x)
}

func min(arr ...int) int{
    if len(arr) == 0 {
        return 0
    }
    min := arr[0]
    for _,v:=range arr{
        if v< min {
            min = v
        }
    }
    return  min
}