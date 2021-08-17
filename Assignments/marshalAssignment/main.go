package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//get Data From Url
	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:8085/inventory")

	if err != nil {
		panic(err)
	}
	newHotel := Hotel{Source: "XYZ", Name: "ABC", RoomType: "SMALL", Price: 100}
	newHotel2 := Hotel{Source: "PQR", Name: "LMN", RoomType: "LARGE", Price: 300}
	//append new Hotels
	hotels = append(hotels, newHotel, newHotel2)
	//get data in json format
	// hotelData, err := json.Marshal(hotels)
	hotelData, err := json.MarshalIndent(hotels, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(hotelData))
}
