package hotelrepo

import "sort"

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

type Hotels []Hotel
type ByRoomType []Hotel

type lessFunc func(p1, p2 *Hotel) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	hotels []Hotel
	less   []lessFunc
}

func NewHotel(src, name, roomtype string, price float32) Hotel {
	return Hotel{
		Source:   src,
		Name:     name,
		RoomType: roomtype,
		Price:    price,
	}
}

func (h ByRoomType) Len() int           { return len(h) }
func (h ByRoomType) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ByRoomType) Less(i, j int) bool { return h[i].RoomType < h[j].RoomType }

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.hotels)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.hotels[i], ms.hotels[j] = ms.hotels[j], ms.hotels[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.hotels[i], &ms.hotels[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(hotels []Hotel) {
	ms.hotels = hotels
	sort.Sort(ms)
}

func (hotels Hotels) FilterByFn(fn func(h *Hotel) bool) Hotels {
	var filteredList Hotels
	//reference semantics: avoid copy
	for i := range hotels {
		if fn(&hotels[i]) {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}
