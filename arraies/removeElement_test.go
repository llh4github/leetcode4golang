package arraies

import "testing"

func TestRemoveElement(t *testing.T) {

	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	rs := removeElement(a, 2)
	t.Log(rs)

}

// 移除元素
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
func removeElement(nums []int, val int) int {
	left, right := 0, 0
	n := len(nums)

	for right < n {
		if nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return left
}
