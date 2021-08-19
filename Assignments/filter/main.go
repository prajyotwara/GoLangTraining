package main

import "fmt"

func main() {
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
		return h.RoomType == `"LARGE" || "X-LARGE"` && h.Price < 130
		// return (h.RoomType == "LARGE" || h.RoomType == "X-LARGE") && h.Price < 130
	})
	fmt.Println(customFilteredHotels)
}
