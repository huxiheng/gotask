package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // k 用来标记唯一元素的个数
    k := 1

    // 从第二个元素开始遍历
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[k] = nums[i] // 把当前不重复的元素放到 nums[k] 位置
            k++
        }
    }

    return k
}

func main() {
    nums1 := []int{1, 1, 2}
    nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

    fmt.Println(removeDuplicates(nums1)) // 输出: 2
    fmt.Println(removeDuplicates(nums2)) // 输出: 5
}
