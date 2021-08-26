package main

import (
	"fmt"
	"time"
)

func SelectDemo() {
	// ReadFromMultipleChannels()
	ReadFromMultipleChannelsWithSelect()
}
func ReadFromMultipleChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//Slow producer
	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 800)
			ch1 <- i
		}
	}()

	//Fast producer
	go func() {
		defer close(ch2)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 150)
			ch2 <- i * 100
		}
	}()

	for {

		out, ok := <-ch1
		if ok {
			fmt.Println("Reading from CH1: ", out)
		}

		out, ok = <-ch2
		if ok {
			fmt.Println("Reading from CH2: ", out)
		}

		//fmt.Println("Continue looping")
	}

}
func ReadFromMultipleChannelsWithSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//Slow producer
	go func() {
		defer close(ch1)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			ch1 <- i
		}
	}()

	//Fast producer
	go func() {
		defer close(ch2)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			ch2 <- i * 100
		}
	}()

	//Select allows to consume values from fast producer fast and also
	//allow fast producer to produce values faster and not block it while sending values on channel
	for ch1 != nil || ch2 != nil {
		select {
		case out, ok := <-ch1:
			if ok {
				fmt.Println("Reading from CH1: ", out)
			} else {
				ch1 = nil
			}
		case out, ok := <-ch2:
			if ok {
				fmt.Println("Reading from CH2: ", out)
			} else {
				ch2 = nil
			}
		default:
			//This means that there is no value on either of the channel
			//So intead of agressively looping just sleep for few milli seconds and try again
			//To prevent for loop to hog the cpu
			time.Sleep(time.Millisecond * 100)
		}
		//un comment to see for loop activity
		//fmt.Println("continue loop")
	}

}
