package arraies

import "testing"

func TestMoveZeroes(t *testing.T) {

	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	t.Log(a)
}

// 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	left, right := 0, 0
	n := len(nums)
	for right < n {
		if nums[right] != 0 {

			// temp := nums[left]
			// nums[left] = nums[right]
			// nums[right] = temp
			// 简化形式
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
