package main

import "fmt"

func main() {
	//4) найти разницу между максимальным и минимальными элементами массива
	var n int
	fmt.Scanf("%d", &n)
	arr := [20]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	min := arr[0]
	max := min
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Printf("Максимальная сумма: %d \n", max)
	fmt.Printf("Минимальная сумма: %d \n", min)
	fmt.Println("Разница между максимальной и минимальной: ", max-min)
}
