package main

import "fmt"

type student struct{
	name string
	id int
	section string
}

func print_student_data(s student){

	fmt.Println("Name: ",s.name)
	fmt.Printf("ID: %d",s.id)	
	fmt.Println()
	fmt.Println("Section: ",s.section)

}

func main(){

	var s student
	s.name="Hasnain"
	s.id = 1
	s.section="CS-D"
	print_student_data(s)

}

