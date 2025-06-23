package main

import (
    "fmt"
    "sort"
)

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    // 按照区间的起始位置进行排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // 初始化结果数组
    result := [][]int{intervals[0]}

    // 遍历剩下的区间
    for i := 1; i < len(intervals); i++ {
        // 获取当前区间和上一个合并区间
        last := result[len(result)-1]
        current := intervals[i]

        // 判断是否有重叠
        if current[0] <= last[1] {
            // 合并区间
            last[1] = max(last[1], current[1])
        } else {
            // 没有重叠，直接加入结果
            result = append(result, current)
        }
    }

    return result
}

// 辅助函数，返回两个数中的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    intervals := [][]int{
        {1, 3},
        {2, 6},
        {8, 10},
        {15, 18},
    }

    result := merge(intervals)
    fmt.Println(result) // 输出: [[1, 6], [8, 10], [15, 18]]
}
