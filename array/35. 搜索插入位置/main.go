package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}
	fmt.Println(searchInsert(nums, 4))
}

func searchInsert(nums []int, target int) int {
	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	var (
		low  = 0
		high = len(nums) - 1
		mid  = (low + high + 1) / 2
	)

	for low < mid {
		if nums[mid] >= target && nums[mid-1] < target {
			return mid
		}
		if nums[mid] < target {
			low = mid
		} else if nums[mid] > target {
			high = mid
		}
		mid = (low + high + 1) / 2
	}
	return -1
}
