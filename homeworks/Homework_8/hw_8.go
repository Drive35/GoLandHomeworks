package main

import (
	"fmt"
	"math"
)

func GetMinMax(a []int) (int, int) {
	min := a[0]
	max := min
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		} else if a[i] < min {
			min = a[i]
		}
	}
	return min, max
}
func GetAverage(b []int) int {
	sumi := 0
	average := 0
	for i := 0; i < len(b); i++ {
		sumi += b[i]
	}
	average = sumi / len(b)
	return average
}
func GetEven(c []int) []int {
	even := []int{}
	for i := 0; i < len(c); i++ {
		if c[i]%2 == 0 {
			even = append(even, c[i])
		}
	}
	return even
}
func GetOdd(d []int) []int {
	odd := []int{}
	for i := 0; i < len(d); i++ {
		if d[i]%2 != 0 {
			odd = append(odd, d[i])
		}
	}
	return odd
}
func GetPow(f []int) []int {
	var h int
	fmt.Println("На какое число возвести степень: ")
	fmt.Scanf("%d \n", &h)
	Pow := []int{}
	for i := 0; i < len(f); i++ {
		P := math.Pow(float64(f[i]), float64(h))
		Pow = append(Pow, int(P))
	}
	return Pow
}
func NumberCheck(e []int) bool {
	var g int
	fmt.Println("Введите число для проверки: ")
	fmt.Scanf("%d", &g)
	numcheck := false
	for i := 0; i < len(e); i++ {
		if g == e[i] {
			numcheck = true
			break
		}
	}
	return numcheck
}
func printInfo(j []int) {
	min, max := GetMinMax(j)
	average := GetAverage(j)
	even := GetEven(j)
	odd := GetOdd(j)
	pow := GetPow(j)
	numcheck := NumberCheck(j)
	fmt.Println("Минимальное число: ", min)
	fmt.Println("Максимальное число: ", max)
	fmt.Println("Среднее число: ", average)
	fmt.Println("Четные числa: ", even)
	fmt.Println("Не четные числа: ", odd)
	fmt.Println("Массив в степени: ", pow)
	fmt.Println("Проверка числа: ", numcheck)
}
func main() {
	arr := []int{11, 2, 7, 3, 12, 5, 10}
	printInfo(arr)
}
