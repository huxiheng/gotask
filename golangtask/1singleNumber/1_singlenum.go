package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
    var nums = []int{2, 2, 3, 3, 5}
    num := singleNumber(nums)

    fmt.Println("只出现一个数字：",num)
}
