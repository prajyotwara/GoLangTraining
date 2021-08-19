package main

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
type HotelPrice float32

type UUID string

type Product struct {
	ID   UUID
	Name string
}

func NewHotel(source, name, roomType string, price float32) Hotel {
	return Hotel{
		Source:   source,
		Name:     name,
		RoomType: roomType,
		Price:    price,
	}
}

//function
func GetPrice(h Hotel) HotelPrice {
	return HotelPrice(h.Price)
}

//method
func (h Hotel) GetPrice() HotelPrice {
	return HotelPrice(h.Price)
}

//function
func SetPrice(h Hotel, newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//method withn value reciever
func (h Hotel) SetPrice(newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//function
func SetHotelPrice(h *Hotel, newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//method with pointer receiver
func (h *Hotel) SetHotelPrice(newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//[]Hotel ---->filter---> []Hotel

//function
func FilterHotelByRoomType(hotelList []Hotel) []Hotel {
	return []Hotel{}
}

type Hotels []Hotel

type HotelFilterFn func(*Hotel) bool

//method with copy semantics
func (hotels Hotels) FilterHotelByRoomTypeWithCopy(roomType string) Hotels {
	var filteredList Hotels

	//copy semantics
	for _, h := range hotels {
		if h.RoomType == roomType {
			filteredList = append(filteredList, h)
		}
	}
	return filteredList
}

func (hotels Hotels) FilterHotelByRoomType(roomType string) Hotels {
	var filteredList Hotels

	//reference semantics
	for i := range hotels {
		if hotels[i].RoomType == roomType {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}

//higher order function as it is accepting func as a value
func (hotels Hotels) FilterByFn(fn func(h *Hotel) bool) Hotels {
	var filteredList Hotels

	for i := range hotels {
		if fn(&hotels[i]) {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}

func (hotels Hotels) FilterBy(fn HotelFilterFn) Hotels {
	var filteredList Hotels

	for i := range hotels {
		if fn(&hotels[i]) {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}
