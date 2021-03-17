package main

import "fmt"

func getFullName(s Student) {
	fmt.Println(s.FirstName, s.LastName)
}
func getAverageMark(a Student) {
	sumi := 0
	average := 0
	for i := 0; i < len(a.Marks); i++ {
		sumi += a.Marks[i]
	}
	average = sumi / len(a.Marks)
	fmt.Println(average)
}

type Student struct {
	FirstName string
	LastName  string
	Marks     []int
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
	students := []Student{s1, s2}
	for _, v := range students {
		getFullName(v)
		getAverageMark(v)
	}
}
