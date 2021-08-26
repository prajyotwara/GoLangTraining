package main

import (
	"encoding/json"
	"fmt"
	"hotelapp/hotelrepo"
	"sort"
)

func main() {
	SortByUsingSortWithSlice()
	SortByUsingSortWithSort()
	SortByRoomTypeAndPrice()
}

func GetHotels() hotelrepo.Hotels {
	var hotels hotelrepo.Hotels
	hotels1, err := hotelrepo.FetchHotelsFromUrlUsingJsonDecoder("http://localhost:3000/inventory")
	if err != nil {
		panic(err)
	}
	hotels = append(hotels, hotels1...)
	hotels2, err := hotelrepo.FetchHotelsFromUrlUsingJsonDecoder("http://localhost:4000/inventory")

	if err != nil {
		panic(err)
	}

	hotels = append(hotels, hotels2...)
	hotels3 := hotelrepo.Hotels{
		{Source: "XYZ", Name: "ABC", RoomType: "SMALL", Price: 100},
		{Source: "PQR", Name: "LMN", RoomType: "LARGE", Price: 300},
		{Source: "PQR", Name: "LMN", RoomType: "X-LARGE", Price: 200},
		{Source: "PQR", Name: "LMN", RoomType: "SMALL", Price: 150},
	}
	hotels = append(hotels, hotels3...)
	return hotels
}

//using closure
func SortByUsingSortWithSlice() {
	hotels := GetHotels()
	sort.Slice(hotels, func(i, j int) bool { return hotels[i].RoomType < hotels[j].RoomType })

	hotelData, err := json.MarshalIndent(hotels, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println("Sorting by using sort.Sort() : ")
	fmt.Println(string(hotelData))
	fmt.Println("==================================================================================")
}

func SortByUsingSortWithSort() {
	hotels := GetHotels()
	sort.Sort(hotelrepo.ByRoomType(hotels))

	hotelData, err := json.MarshalIndent(hotels, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println("Sorting by using sort.Slice()")
	fmt.Println(string(hotelData))
	fmt.Println("==================================================================================")

}

func SortByRoomTypeAndPrice() {
	// Closures that order the Change structure.
	roomType := func(h1, h2 *hotelrepo.Hotel) bool {
		return h1.RoomType < h2.RoomType
	}
	price := func(h1, h2 *hotelrepo.Hotel) bool {
		return h1.Price < h2.Price
	}
	hotels := GetHotels()

	hotelrepo.OrderedBy(roomType, price).Sort(hotels)
	hotelData, err := json.MarshalIndent(hotels, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Println("Sorting by roomType and Price")
	fmt.Println(string(hotelData))
	fmt.Println("==================================================================================")
}

func FilterHotels() {
	hotels := GetHotels()

	var roomFilter []string
	type PriceFilter struct {
		Price     float32
		Condition string
	}
	var priceFilter []PriceFilter
	roomFilter = append(roomFilter, "SMALL")
	priceFilter = append(priceFilter, PriceFilter{Price: 200, Condition: "<"})

	customFilteredHotels := hotels.FilterByFn(func(h *hotelrepo.Hotel) bool {
		if roomFilter != nil {
			if priceFilter.Price != 0 {

			}
		}
		return (h.RoomType == "LARGE" || h.RoomType == "X-LARGE") && h.Price < 130
	})
}
