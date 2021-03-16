package main

import "fmt"

func GetB(c []int) int {
	a := 0
	b := 0
	for i := 0; i < len(c); i++ {
		if c[i] < 1 {
			a = 0
		} else {
			a = a + 1
		}
		if a > b {
			b = a
		}
	}
	return b
}
func printB(d []int) {
	number := GetB(d)
	fmt.Println("Длина самой продолжительной оттепели: ", number)
}
func main() {
	arr1 := []int{-20, 30, -40, 50, 10, -10}
	arr2 := []int{10, 20, 30, 1, -10, 1, 2, 3}
	arr3 := []int{-10, 0, -10, 0, -10}
	printB(arr1)
	printB(arr2)
	printB(arr3)
}

