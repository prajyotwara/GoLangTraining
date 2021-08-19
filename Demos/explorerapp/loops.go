package main

import "fmt"

func GoLangForLoopForms() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While loop like behavior
	count := 0
	for count < 10 {
		count += 1
		fmt.Println(count)
	}

	count2 := 0
	//While loop like behavior
	for {
		count2 += 1
		if count2 > 10 {
			break
		}
		fmt.Println(count2)
	}

}
