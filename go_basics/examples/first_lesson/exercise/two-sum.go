package main

import "fmt"

func main() {
	target := 9
	arr := [4]int{2, 7, 8, 9}
	indexes := twoSum(arr[:], target)
	fmt.Printf("index1: %d, index2: %d\n", indexes[0], indexes[1])
}

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
