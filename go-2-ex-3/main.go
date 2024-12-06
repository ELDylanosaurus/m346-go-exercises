package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Einführung in die Programmierung",
		117: "Algorithmen und Datenstrukturen",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117) 

	modules[220] = "Webentwicklung mit Go" 

	modules[346] = "DevOps und Cloud-Management" 

	fmt.Println(modules)
}
