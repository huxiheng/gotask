package main

import (
    "fmt"
)

func plusOne(digits []int) []int {
    n := len(digits)

    // 从最低位开始处理
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        // 如果是 9，则变成 0，继续进位
        digits[i] = 0
    }

    // 如果所有位都是 9，例如 [9,9,9]，需要扩展一位
    result := make([]int, n+1)
    result[0] = 1
    return result
}

func main() {
    fmt.Println(plusOne([]int{1, 2, 3})) // 输出: [1 2 4]
    fmt.Println(plusOne([]int{9}))       // 输出: [1 0]
    fmt.Println(plusOne([]int{9, 9, 9})) // 输出: [1 0 0 0]
}
