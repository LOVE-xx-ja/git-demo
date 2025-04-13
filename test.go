package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, j*i)
		}
		fmt.Println()
	}

}
