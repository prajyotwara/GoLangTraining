package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
{
	source: "gold",
	name: "Hotel-AB44",
	room-type: "LARGE",
	price: "200"
}
*/
//json tags tell the json Marshal & unmarshall methods how to convert the fields to & from json
//like RoomType would be Marshalled to room-type
//and Price would be marshalled to price and since it is string in json and float32 in code it needs conversion

type Hotel struct {
	Source   string  `json:"source"`
	Name     string  `json:"name"`
	RoomType string  `json:"room-type"`
	Price    float32 `json:"price,string"`
}

// func NewHotel() Hotel {

// }

func FetchHotelsFromUrl(url string) ([]Hotel, error) {

	client := http.Client{}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var hotels []Hotel

	err = json.Unmarshal(databytes, &hotels)

	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func FetchHotelsFromUrlUsingJsonDecoder(url string) ([]Hotel, error) {

	client := http.Client{}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var hotels []Hotel

	json.NewDecoder(response.Body).Decode(&hotels)
	return hotels, nil
}
