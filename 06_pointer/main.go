package main

import "fmt"

type Student struct {
	ID int
	Name string
	GPA float32
}

func (s *Student) SetStudent(id int, name string, gpa float32) {
	s.ID = id
	s.Name = name
	s.GPA = gpa
}

func (s *Student) Graduate() {
	s.Name = s.Name + " S.TI."
}

func main() {
	numberA := 5
	numberB := &numberA

	fmt.Println(numberA, numberB, *numberB)
	*numberB = 7
	fmt.Println(numberA, numberB, *numberB)

	var numberC int = 5
	var numberD *int = &numberC
	var numberE int = numberC

	fmt.Println(numberC, numberD, *numberD, numberE)

	numberC = 20
	fmt.Println(numberC, numberD, *numberD, numberE)
	fmt.Printf("\n\n")

	number := 5
	fmt.Println("Nilai Awal : ", number)
	fmt.Println("Alamat Memory: ", &number)

	change(&number, 100)
	fmt.Println("Nilai Baru : ", number)
	fmt.Println("Alamat Memory: ", &number)
	
	ary := Student{}
	ary.SetStudent(1, "Ary Setiyawan", 3.69)

	fmt.Println(ary)
	ary.Graduate()
	fmt.Println(ary)
	
}

func change(oldValue *int, newValue int) {
	*oldValue = newValue
	fmt.Println("Nilai Dalam Function: ", *oldValue)
	fmt.Println("Alamat Memory Dalam Function: ", oldValue)
}