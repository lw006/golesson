package main //必须

import (
	"fmt"

	"./calc"
	"./calc/test"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)

	fmt.Println("r = ", calc.Minus(10, 5))
	test.Test1()
}
