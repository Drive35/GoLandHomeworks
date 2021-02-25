package main

import "fmt"

func main() {
	//2) ряд фибоначчи
	//вводится какое либо число и небходимо посчитать число фибоначи до этого числа
	//n=10
	//0 1 1 2 3 5 8 13

	var a, b, c, n int
	fmt.Scanf("%d", &n)
	a = 1
	b = 2
	fmt.Printf("%d %d ", a, b)
	for i := 3; i <= n; i++ {
		fmt.Printf("%d ", a+b)
		c = a
		a = b
		b = c + a
	}
	fmt.Printf("\n")
}
