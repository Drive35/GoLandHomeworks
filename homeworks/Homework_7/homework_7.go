package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := [15]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d \n", &a)
		arr[i] = a
	}
	var b int
	fmt.Println("Введите число")
	fmt.Scanf("%d \n", &b)
	c := []int{}
	for i := 0; i < n; i++ {
		if b != arr[i] {
			c = append(c, arr[i])
		}
	}
	fmt.Println(c)
}
