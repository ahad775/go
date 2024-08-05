package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	RollNo    int
}

func main() {

	student := Student{
		FirstName: "ahad",
		LastName:  "naz",
		RollNo:    2,
	}
	fmt.Println(student.FirstName)
}