package twosum

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
*/

func Run() {
	arr := []int{2, 7, 11, 15}
	target := 9
	res1 := solution1(arr, target)
	fmt.Println(res1)

	res2 := solution2(arr, target)
	fmt.Println(res2)
}

func solution1(arr []int, target int) (res []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i]+arr[j] == target {
				res = append(res, i)
				res = append(res, j)
				return res
			}
		}
	}
	return res
}

func solution2(arr []int, target int) (res []int) {
	m := make(map[int]int)
	for k, v := range arr {
		m[v] = k // {2:0, 7:1 }
	}

	for i, j := range m {
		another := target - i
		if anotherKey, ok := m[another]; ok {
			return []int{j, anotherKey}
		}
	}
	return res
}
