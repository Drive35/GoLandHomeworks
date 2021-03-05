package main

import "fmt"

func main() {
	//2)
	//вводятся числа в массив там могут быть и отрицательные и положительные и вам необходимо
	//посчитать сумму элементов после первого отрицательного числа
	//
	//1 2 3 -1 2 3 4
	//Ответ:9

	var n, b int
	fmt.Scanf("%d", &n)
	arr := [60]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			b = i
			break
		}
	}
	sumd := 0
	for i := b + 1; i < n; i++ {
		sumd = arr[i] + sumd
	}
	fmt.Println(sumd)
}
