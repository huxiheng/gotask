package main

import (
    "fmt"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for i := 0; i < len(strs[0]); i++ {
        ch := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != ch {
                return strs[0][:i]
            }
        }
    }

    return strs[0]
}

func main() {
    // 示例 1
    strs1 := []string{"flower", "flow", "flight"}
    fmt.Println("最长公共前缀：", longestCommonPrefix(strs1)) // 输出："fl"

    // 示例 2
    strs2 := []string{"dog", "racecar", "car"}
    fmt.Println("最长公共前缀：", longestCommonPrefix(strs2)) // 输出：""
}
