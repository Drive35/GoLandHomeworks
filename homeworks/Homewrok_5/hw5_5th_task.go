package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := [100]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	J := [100]int{}
	K := [100]int{}
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			J[i] = arr[i]
		} else {
			K[i] = arr[i]
		}
	}
	fmt.Println("Четный массив = ", J)
	fmt.Println("Не четный массив = ", K)
}
