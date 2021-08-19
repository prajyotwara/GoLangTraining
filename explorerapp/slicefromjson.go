package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
