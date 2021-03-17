package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := []int{}
	for i := 0; i < n; i++ {
		var g int
		fmt.Scanf("%d", &g)
		arr[i] = g
	}
	for i := 0; i < n; i++ {
		h := true
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] && i != j {
				h = false
				break
			}
		}
		if h == true {
			fmt.Println(arr[i])
		}
	}
}
