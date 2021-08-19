package counters

import "fmt"

type Counter interface {
	Increment()
	Decrement()
	GetValue() int64
}

func CounterOperations(counter Counter) {
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.GetValue())
}
