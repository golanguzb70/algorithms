package main

import "fmt"

/*
	Array in Golang
	Insertion: O(n)
	Deletion: O(n)
	Update: O(1)
	Access: O(1)
*/

type MyArray struct {
	Nums []int
}

func NewNumsArray(in []int) *MyArray {
	return &MyArray{
		Nums: in,
	}
}

func (arr *MyArray) InsertNumber(index int, value int) {
	if len(arr.Nums) <= index {
		arr.Nums = append(arr.Nums, value)
	} else {
		lastValue := 0
		for i := range arr.Nums {
			if i == index {
				lastValue = arr.Nums[i]
				arr.Nums[i] = value
			} else if i > index {
				temp := arr.Nums[i]
				arr.Nums[i] = lastValue
				lastValue = temp
			}
		}
		arr.Nums = append(arr.Nums, lastValue)
	}
}

func (arr *MyArray) DeleteIndex(index int) {
	if len(arr.Nums) <= index {
		return
	}

	for i := range arr.Nums {
		if i >= index && i != len(arr.Nums)-1 {
			arr.Nums[i] = arr.Nums[i+1]
		}
	}

	arr.Nums[len(arr.Nums)-1] = 0
}




func main() {
	fmt.Println("Array in Golang")
}
