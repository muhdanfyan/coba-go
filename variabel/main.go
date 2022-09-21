package main

import "fmt"

func main() {
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	// boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
}
