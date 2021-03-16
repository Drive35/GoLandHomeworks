package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray(a []int) []int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	arr := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(10)
		arr = append(arr, k)
	}
	return arr
}
func getSumC(b []int) int {
	c := createArray(b)
	sumc := 0
	for i := 0; i < len(c); i++ {
		sumc = sumc + c[i]
	}
	return sumc
}
func printC(d []int) {
	maxnumber := getSumC(d)
	fmt.Println("Максимальное число ягод: =", maxnumber)
}
func main() {
	arr := []int{}
	printC(arr)
}
