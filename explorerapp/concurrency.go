package main

import (
	"fmt"
	"time"
)

func GoKeywordDemo() {
	SomeSimpleFunction()
	go func() {
		fmt.Println("Inside Anonymous func(){}")
	}()
	sfnValue := func() {
		fmt.Println("Inside Simple fn defined as function value")
	}
	go sfnValue()
	time.Sleep(time.Second * 1)
}

func SomeSimpleFunction() {
	fmt.Println("Inside Simple Function")
}
