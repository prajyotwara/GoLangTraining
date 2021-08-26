package hotelrepo

import (
	"encoding/json"
	"net/http"
)

func FetchHotelsFromUrlUsingJsonDecoder(url string) (Hotels, error) {

	client := http.Client{}

	response, err := client.Get(url)

	//if not nil then there is some error
	if err != nil {
		//unhappy path, just return in this case. Should also log the error.
		return nil, err
	}

	defer response.Body.Close()

	var hotels Hotels

	json.NewDecoder(response.Body).Decode(&hotels)

	return hotels, nil

}
