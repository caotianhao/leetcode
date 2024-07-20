package main

import "fmt"

func isConsecutive(nums []int) bool {
	m := map[int]struct{}{}
	min, l := 100001, len(nums)
	for _, v := range nums {
		if v < min {
			min = v
		}
		m[v] = struct{}{}
	}
	for i := min; i <= min+l-1; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isConsecutive([]int{5, 6, 8, 8}))
}