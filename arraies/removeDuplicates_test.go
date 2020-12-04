package arraies

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	rs := removeDuplicates(a)
	t.Log(rs)
}

// 删除排序数组中的重复项
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
func removeDuplicates(nums []int) int {
	left, right := 0, 0
	n := len(nums)

	for right < n {
		if nums[right] != nums[left] {
			left++
			nums[left], nums[right] = nums[right], nums[left]
		}
		right++
	}

	return left + 1 // 数组长度需要+1
}
