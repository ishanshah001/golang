package main

import "fmt"

type Student struct {
	Name string
	Age  int
	GPA  float64
}

// Value receiver: just prints
func (s Student) PrintDetails() {
	fmt.Printf("Name: %s, Age: %d, GPA: %.2f\n", s.Name, s.Age, s.GPA)
}

// Pointer receiver: modifies original
func (s *Student) UpdateGPA(newGPA float64) {
	s.GPA = newGPA
}

// function
func PrintStudent(s Student) {
	fmt.Printf("Name: %s, Age: %d, GPA: %.2f\n", s.Name, s.Age, s.GPA)
}

func UpdateAge(s *Student, newAge int) {
	s.Age = newAge
}

func main() {
	s := Student{Name: "Ishan", Age: 23, GPA: 3.5}
	s.PrintDetails()

	s.UpdateGPA(3.9)
	s.PrintDetails()
	PrintStudent(s)
	UpdateAge(&s, 24)
	PrintStudent(s)

}
