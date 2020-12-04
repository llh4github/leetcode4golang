package arraies

import "testing"

// 颜色分类
// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，
// 使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

func TestSortColors1(t *testing.T) {
	a := []int{2, 0, 2, 1, 1, 0}
	b := []int{2, 0}
	sortColors1(a)
	sortColors1(b)
	t.Log(a)
	t.Log(b)
}

// 解法一：计数替换法
// 因为只需对三个数排序，可以先对各个数进行计数，
// 再根据个数重写数组内容。
func sortColors1(nums []int) {
	c0, c1, c2 := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			c0++
		case 1:
			c1++
		case 2:
			c2++
		default:
			println("数字不合法！")
		}
	}

	for x := 0; x < c0; x++ {
		nums[x] = 0
	}
	for x := 0; x < c1; x++ {
		nums[c0+x] = 1
	}
	for x := 0; x < c2; x++ {
		nums[c0+c1+x] = 2
	}
}
