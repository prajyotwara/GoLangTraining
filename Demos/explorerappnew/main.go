package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ChannelBasics()
	// ChannelWorkers()
	// SelectDemo()
	PipelineDemo()
}

func ChannelBasics() {
	//unbuffered channel
	// ch := make(chan int64)

	//buffered channel
	ch := make(chan int64, 3)

	//Sender sending 3 values on a channel
	go func() {
		defer close(ch)
		ch <- 5
		ch <- 10
		ch <- 15
	}()

	// out1 := <-ch
	// fmt.Println(out1)
	// out2 := <-ch
	// fmt.Println(out2)
	// out3 := <-ch
	// fmt.Println(out3)
	time.Sleep(time.Second * 5)

	//Receiver reading values from the channel
	for out := range ch {
		fmt.Println(out)
	}
}

//Channels are syncronised
//Multiple sender can send into a channel
//Multiple receiver can receive from a channel
//Sending receiving on a close channel is bad
func ChannelWorkers() {
	ch := make(chan int64)

	go func() {
		defer close(ch)
		ch <- 5
		ch <- 12
		ch <- 18
		ch <- 21
		ch <- 11
		ch <- 22
		ch <- 13
		ch <- 10
	}()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for out := range ch {
				fmt.Println("Worker number : ", num, " Working on Item No : ", out)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Done working on go work routines")
}
