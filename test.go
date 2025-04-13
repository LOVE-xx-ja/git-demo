package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//冒泡排序
	var arr [10]int = [...]int{43, 32, 5, 7, 8, 98, 65, 432, 32, 2}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	//遍历
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}

	//99乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, j*i)
		}
		fmt.Println()
	}
	fmt.Println("github pull test")

}
