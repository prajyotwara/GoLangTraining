package server

import (
	"encoding/json"
	"net/http"
	"time"
)

/*
{
	source: "gold",
	name: "Hotel-AB44",
	room-type: "LARGE",
	price: "200"
}
*/
type Hotel struct {
	Source   string  `json:"source"`
	Name     string  `json:"name"`
	RoomType string  `json:"room-type"`
	Price    float32 `json:"price,string"`
}

func FetchDataFromUrl(url string) (<-chan Hotel, error) {
	hotel := make(chan Hotel)
	client := http.Client{}
	response, err := client.Get(url)

	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second * 20)
	defer response.Body.Close()
	var hotels []Hotel
	json.NewDecoder(response.Body).Decode(&hotels)

	go func() {
		defer close(hotel)
		for _, v := range hotels {
			hotel <- v
		}
	}()
	return hotel, nil
}
