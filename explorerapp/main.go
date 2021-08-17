package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Day1()
	// Day2()
	Day3()

}

func Day1() {
	// fmt.Println("Hello Go")
	// VariablesAndPointer()

	// email := services.NewEmail("prajyotw@cybage.com", "GoLang Training", "Details about the training .....")

	// services.SendEmail(email)

	// practices.Practice()

}
func Day2() {
	// GoLangForLoopForms()
	//GoKeywordDemo()
	// ConcurrencyWithWaitGroupExample()
	// LinkedListOps()
}
func Day3() {
	// GoArrayAndSlices()
	// GoSlices()
	// GoSlicesRange()
	// hotels, err := FetchHotelsFromUrl("http://localhost:8085/inventory")
	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:8085/inventory")

	if err != nil {
		panic(err)
	}
	newHotel := Hotel{Source: "XYZ", Name: "ABC", RoomType: "SMALL", Price: 100}
	newHotel2 := Hotel{Source: "PQR", Name: "LMN", RoomType: "LARGE", Price: 300}
	hotels = append(hotels, newHotel, newHotel2)
	// hotelData, err := json.Marshal(hotels)
	hotelData, err := json.MarshalIndent(hotels, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(hotelData))
}
