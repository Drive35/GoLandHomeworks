package main

import "fmt"

func main() {
	//2) найти сумму положительных элементов в массиве

	var n int
	fmt.Scanf("%d", &n)
	arr := [25]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
		if arr[i] > 0 {
			fmt.Println(arr[i])
		}
	}
}

