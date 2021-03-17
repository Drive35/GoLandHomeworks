package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateH(rows, columns, randomRange int) [][]int {
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
func Getsumi(a [][]int) int {
	sumi := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			sumi = sumi + a[i][j]
		}
	}
	return sumi
}
func getElemIndex(a [][]int) (int, int)  {
	mm := 0
	nx := 0
	sump :=Getsumi(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if sump > mm {
				mm = sump
				nx = i
			}
		}
	}
	return mm, nx
}

func FuncT(a [][]int) {
	sump := Getsumi(a)
	maxelem, nestindex := getElemIndex(a)
	fmt.Println("Сумма элементов:= ", sump)
	fmt.Println("Максимальная сумма массива = ", maxelem)
	fmt.Println("Строка максимального массива = ", nestindex)
}
func main() {
	arr := CreateH(10, 6, 200)
	for _, v := range arr {
		fmt.Println(v)
	}
	FuncT(arr)
}

