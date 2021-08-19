package hotelsrepo

import (
	"encoding/json"
	"net/http"
)

type ApiHotelRepo struct {
	ApiUrl string
}

func NewApiHotelRepo(url string) ApiHotelRepo {
	return ApiHotelRepo{ApiUrl: url}
}

func (hr *ApiHotelRepo) GetAllHotels() ([]Hotel, error) {
	client := http.Client{}
	response, err := client.Get(hr.ApiUrl)

	if err != nil {
		return []Hotel{}, err
	}

	defer response.Body.Close()

	var hotels []Hotel

	json.NewDecoder(response.Body).Decode(&hotels)

	return hotels, nil
}
