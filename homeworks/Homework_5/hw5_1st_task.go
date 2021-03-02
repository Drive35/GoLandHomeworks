package main

import "fmt"

func main () {
	//1) найти количество четных и нечетных элементов в массиве

	var n int
	fmt.Scanf("%d", &n)
	arr := [50]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	b := 0
	c := 0
	for j := 0; j < n; j++ {
		d := arr[j]
		if d%2 == 0 {
			b += 1
		} else {
			c += 1
		}
	}
	fmt.Printf("Количество четных = %d \n", b)
	fmt.Printf("Количество нечетных = %d \n", c)
}
