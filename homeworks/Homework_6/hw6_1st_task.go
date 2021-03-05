package main

import "fmt"

func main () {
	//1) вводятся любые числа в массив
	//и среди них необходимо найти те числе которые являются четными и положить их индексы
	//в другой массив и вывести
	//
	//4 9 2 3 4 5 1
	//0 1 2 3 4 5 6
	//
	//Ответ:0 2 4

	var n int
	fmt.Scanf("%d", &n)
	arr := [70]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	c := [70]int{}
	i2 := 0
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			c[i2] = i
			i2++
		}
	}
	fmt.Println(c)
}
