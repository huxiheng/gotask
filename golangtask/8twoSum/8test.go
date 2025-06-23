package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 创建一个哈希表用于存储数值和其对应的索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算当前数和目标值的差
		complement := target - num

		// 如果差值已经在哈希表中，返回结果
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		// 如果差值不在哈希表中，将当前数值和索引存入哈希表
		numMap[num] = i
	}

	// 如果没有找到，返回一个空数组（题目保证一定有解）
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result) // 输出: [0, 1]
}
