package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateArrA(rows, columns, randomRange int) [][]int {
	rand.Seed(time.Now().UnixNano())
	arr := [][]int{}
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; {
			k := rand.Intn(100)
			a = append(a, k)
		}
		for j := 1; j < columns; j++ {
			l := rand.Intn(randomRange)
			a = append(a, l)
		}
		arr = append(arr, a)
	}
	return arr
}
func main() {
	N := rand.Intn(100)
	b := rand.Intn(1)
	arr := CreateArrA(N, 2, b)
	maxi := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > maxi || arr[i][1] == 1 {
				maxi = arr[i][j]
			} else {
				fmt.Println("Нету мужчин в списке:= ", -1)
			}
			fmt.Println(arr)
			fmt.Println("Количество жильцов:= ", N)
			fmt.Println("Номер самого старшего жителя мужского пола:= ", maxi)
		}
	}
}
