package main

import "fmt"

func main() {
	safe := true
	fmt.Println(safe)

	notSafe := true
	fmt.Println(notSafe)
	//This will also work because Go uses type inference to figure out the data type.

	isFalse := true
	fmt.Println(isFalse)
	//This is an even shorter way and still works.
	//:= declares and initializes the variable.

	name := "Carl"
	isScientist := true
	age := 62
	degree := 5.2
	fmt.Println(name, isScientist, age, degree)

	//*You can not use short declaration at the package scope!
	//The reason for this because at the global scope, you need to use keyword.

}
