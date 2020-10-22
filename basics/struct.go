package main

import "fmt"

type Student struct {
	name    string
	dept    string
	branch  string
	stud_id int
}

func main() {
	var Student1 Student
	var Student2 Student

	Student1.name = "Mythili"
	Student1.dept = "IST"
	Student1.branch = "MCA"
	Student1.stud_id = 6495407

	Student2.name = "Yuvarani"
	Student2.dept = "IST"
	Student2.branch = "MCA"
	Student2.stud_id = 6495700

	fmt.Printf("Student 1 name : %s\n", Student1.name)
	fmt.Printf("Student 1 dept : %s\n", Student1.dept)
	fmt.Printf("Student 1 branch : %s\n", Student1.branch)
	fmt.Printf("Student 1 stud_id : %d\n", Student1.stud_id)

	fmt.Printf("Student 2 name : %s\n", Student2.name)
	fmt.Printf("Student 2 dept : %s\n", Student2.dept)
	fmt.Printf("Student 2 branch : %s\n", Student2.branch)
	fmt.Printf("Student 2 stud_id : %d\n", Student2.stud_id)

}
