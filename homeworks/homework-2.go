package main

import (
	"fmt"
	"math"
)

func main() {
	//1-ая задача
	var a, e, c int
	fmt.Scanf("%d %d %d", &a, &e, &c)
	fmt.Printf("%d%d%d \n", a, e, c)
	p := a + e + c
	fmt.Println(p)

	//2) y = kx + b (k-?,b-?)
	//Ввод
	//3 4
	//5 6
	//x1=3,y1=4
	//x2=5,y2=6

	var x1, x2, y1, y2 int

	k1 := x1 / x1
	b1 := y1 - k1*x1
	k2 := x2 / x2
	b2 := y2 - k2*x2
	fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
	fmt.Printf("Значение k1 = %d \n", k1)
	fmt.Printf("Значение k2 = %d \n", k2)
	fmt.Printf("Значение b1 = %d \n", b1)
	fmt.Printf("Значение b2 = %d \n", b2)

	//3) гипотенуза, вводятся два катета a,b (формула для нахождения гипотенузы прямоугольного треугольника)

	var f, g int
	fmt.Scanf("%d %d", &f, &g)

	h := math.Sqrt(float64(math.Pow(float64(f), 2.0)+math.Pow(float64(g), 2.0)))
	fmt.Println(h)
}
