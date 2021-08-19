package main

import "fmt"

func VariablesAndPointer() {
	var name string = "Prajyot"
	fmt.Println(name)

	role := "Software Engineer"
	fmt.Println(role)

	namePtr := &name

	fmt.Printf("%p	%v \n", namePtr, *namePtr)

	*namePtr = role
	fmt.Printf("%p	%v \n", namePtr, *namePtr)
	fmt.Printf("name = %v	\n", name)

	var aPtr *string
	fmt.Printf("%v ", aPtr)
	//below line will through error as we try to access value <nil> address
	//panic: runtime error: invalid memory address or nil pointer dereference
	// fmt.Printf("%p %v", aPtr, *aPtr)

}
