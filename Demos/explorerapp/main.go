package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Day1()
	// Day2()
	// Day3()
	Day4()

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
func MethodDemo() {
	hotel := NewHotel("gold", "HOTEL-DDF", "SMALL", 20.0)

	p := hotel.GetPrice()
	fmt.Println(p)

	//if use the func instead of method
	//SetPrice(hotel, 55)

	hotel.SetPrice(50)
	fmt.Println("New Price: ", hotel.GetPrice())

	//SetHotelPrice(&hotel, 65)

	hotel.SetHotelPrice(60)
	fmt.Println("After SetHotelPrice: ", hotel.GetPrice())
}

func Day4() {
	// MethodDemo()

	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:8085/inventory")

	if err != nil {
		panic(err)
	}

	fhotels := Hotels(hotels).FilterHotelByRoomType("SMALL")
	fmt.Println(fhotels)

	sHotels := Hotels(hotels).FilterByFn(func(h *Hotel) bool {
		return h.RoomType == "LARGE"
	})
	fmt.Println(sHotels)

	filterFn := func(h *Hotel) bool {
		return h.RoomType == "X-LARGE"
	}
	xlHotels := Hotels(hotels).FilterBy(filterFn)
	fmt.Println(xlHotels)

	customFilteredHotels := Hotels(hotels).FilterByFn(func(h *Hotel) bool {
		return (h.RoomType == "LARGE" || h.RoomType == "X-LARGE") && h.Price < 130
	})
	fmt.Println(customFilteredHotels)
}
