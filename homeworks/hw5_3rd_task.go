package main

import "fmt"

//3) найти количество элементов, которые меньше среднего арифметического элемента и вывести их
//среднее ариф avg = sumi/n

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := [30]int{}
		for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		arr[i] = a
	}
	total := 0
	for i := 0; i < n; i++ {
		total = total + arr[i]
	}
	avg := total/n
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] < avg  {
			fmt.Println("Меньше среднего арифметического", arr[i])
			count += 1
		}
	}
		fmt.Printf("Количество элементов, которые меньше среднего: %d", count)
}
