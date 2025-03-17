package utils

import (
    "fmt"
    "strings"
)

func Clone(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

func FormatNumber(n int) string {
    in := fmt.Sprintf("%d", n)
    out := make([]string, 0, len(in)+(len(in)-1)/3)
    for i, v := range in {
        if i > 0 && (len(in)-i)%3 == 0 {
            out = append(out, ",")
        }
        out = append(out, string(v))
    }
    return strings.Join(out, "")
}