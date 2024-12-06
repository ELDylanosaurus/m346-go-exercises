package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name    string
	Students []Student
}

func main() {
	student1 := Student{FirstName: "Dylan", LastName: "Bellomusto"}
	student2 := Student{FirstName: "Nando", LastName: "Brun"}
	student3 := Student{FirstName: "Leandro", LastName: "Schweggler"}
	student4 := Student{FirstName: "Lisa", LastName: "Eggerschwiler"}
	student5 := Student{FirstName: "Peter", LastName: "Kauffmann"}
	student6 := Student{FirstName: "Sophie", LastName: "Hegi"}

	classA := Class{
		Name: "Class A",
		Students: []Student{student1, student2, student3},
	}
	classB := Class{
		Name: "Class B",
		Students: []Student{student4, student5, student6},
	}

	modules := map[int][]Class{
		104: {classA, classB}, 
		117: {classA},         
		346: {classB},         
	}

	fmt.Println("Classes and their students:")
	fmt.Println(classA)
	fmt.Println(classB)

	fmt.Println("\nModules and attending classes:")
	for module, classes := range modules {
		fmt.Printf("Module %d:\n", module)
		for _, class := range classes {
			fmt.Printf("  %s with students: %v\n", class.Name, class.Students)
		}
	}
}
