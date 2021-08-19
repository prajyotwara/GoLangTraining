package main

import "fmt"

func GoArrayAndSlices() {
	nums := []int32{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(nums))
	fmt.Println(nums)

	numRange := nums[0:4]
	fmt.Printf("Slice 1 of nums %v\n", numRange)

	numRange2 := nums[3:6]
	fmt.Printf("Slice 2 of nums %v\n", numRange2)

}
func GoSlices() {
	numSlice := make([]int32, 0, 10)
	fmt.Println(numSlice)
	fmt.Println(len(numSlice))

	numSlice = append(numSlice, 1, 5, 6, 3, 6)
	fmt.Println(numSlice)
	fmt.Println(len(numSlice))

	numSlice = append(numSlice, 12, 32, 14, 56, 7, 82, 54, 81, 19, 76, 23, 43)
	fmt.Println(numSlice)
	fmt.Println(len(numSlice))
}

func GoSlicesRange() {
	numSlice := make([]int, 0, 10)
	numSlice = append(numSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	for i, v := range numSlice {
		fmt.Println(i, " : ", v)
	}

	for i := 0; i < len(numSlice); i++ {
		fmt.Println(i, " : ", numSlice[i])
	}
	for i := range numSlice {
		numSlice[i] = i
	}
	fmt.Println(numSlice)

	fmt.Println("---------------")

	another := make([]int, 0, 100)
	another = append(another, 1, 2, 4, 10)
	another = append(another, numSlice...)
	fmt.Println(another)
}
