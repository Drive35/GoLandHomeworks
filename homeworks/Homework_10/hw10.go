package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	Marks     []int
}

func (g Student) getFullName() {
	fmt.Println(g.FirstName, g.LastName)
}
func (a Student) getAverageMark() {
	sumi := 0
	average := 0
	for i := 0; i < len(a.Marks); i++ {
		sumi += a.Marks[i]
	}
	average = sumi / len(a.Marks)
	fmt.Println(average)
}

func main() {
	s1 := Student{
		FirstName: "Петр",
		LastName:  "Иванов",
		Marks:     []int{5, 7, 5, 1, 2},
	}
	s2 := Student{
		"Иван",
		"Сидоров",
		[]int{4, 10, 9, 1, 1},
	}
	s1.getFullName()
	s1.getAverageMark()
	s2.getFullName()
	s2.getAverageMark()
}
