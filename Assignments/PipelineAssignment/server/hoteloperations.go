package server

import (
	"sync"
)

func GetHotels() {
	// goldHotels, err := FetchDataFromUrl("http://localhost:3000/inventory")
	// if err != nil {
	// 	panic(err)
	// }
	// silverHotels, err := FetchDataFromUrl("http://localhost:4000/inventory")
	// if err != nil {
	// 	panic(err)
	// }
	// hotels := merge(goldHotels, silverHotels)
	// for i := range hotels {
	// 	fmt.Println(i)
	// }

	urls := []string{"http://localhost:3000/inventory", "http://localhost:4000/inventory", "http://localhost:5000/inventory"}

}

func merge(chans ...<-chan Hotel) <-chan Hotel {
	hotel := make(chan Hotel)

	go func() {
		defer close(hotel)
		var wg sync.WaitGroup
		for _, c := range chans {
			wg.Add(1)
			go func(inCh <-chan Hotel) {
				defer wg.Done()
				for v := range inCh {
					hotel <- v
				}
			}(c)
		}
		wg.Wait()
	}()
	return hotel
}
