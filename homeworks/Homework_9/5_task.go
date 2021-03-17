package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateTDA(rows, columns, randomRange int) [][]int {
	rand.Seed(time.Now().UnixNano())
	arr := [][]int{}
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			k := rand.Intn(randomRange)
			a = append(a, k)
		}
		arr = append(arr, a)
	}
	return arr
}
func main() {
	arr := CreateTDA(10, 20, 15)
	for _, v := range arr {
		fmt.Println(v)
	}
	c := 0
	d := 0
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 5 {
				count += 1
				if count > 3 || count == 3 {
					c = count
					d = i
				}
			}
		}
	}
	fmt.Println("Повторения в массиве = ", c, "", "Строка массива = ", d)
}
