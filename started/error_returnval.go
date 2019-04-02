package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	if ret2, err2 := MySqrt1(9); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
}

func MySqrt(var1 float64) (float64, error) {
	if var1 < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(var1), nil
}

func MySqrt1(var1 float64) (x1 float64, x2 error) {
	if var1 < 0 {
		x1 = float64(math.NaN())
		x2 = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		x1 = math.Sqrt(var1)
	}
	return
}
