package main

import "fmt"

func main() {
	//3)
	//n=5
	//4
	//3
	//10
	//2
	//4
	//На что нужно поделить:2   ->d
	//
	//необходимо вывести количество чисел из массива которые делятся на d без остатка
	//
	//Ответ:4

	var n, b int
	fmt.Scanf("%d \n", &n)
	fmt.Println("На что нужно делить:")
	fmt.Scanf("%d \n", &b)
	arr := [10]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	counter := 0
	for i := 0; i < n; i++ {
		if b == 0 {
			fmt.Println("На 0 нельзя делить")
			break
		} else {
			f := arr[i]
			if f%b == 0 {
				counter += 1
			}
		}
	}
	fmt.Println(arr)
	fmt.Println("Ответ: ", counter)
}
