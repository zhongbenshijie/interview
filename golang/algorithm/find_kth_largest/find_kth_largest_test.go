package find_kth_largest

import (
	"testing"
)

// 求无须数组中第k大的数
// 基于快速排序来实现，因为快速排序从大到小排序，每一轮递归结束时，最后放进去的那个位置就是有特殊意义的

func findKthLargest(nums []int, k int) int {
	return find(0, len(nums)-1, nums, k)
}

func find(left, right int, nums []int, k int) int {
	start, end := left, right

	tmp := nums[left]
	for left < right {
		for left < right && nums[right] <= tmp {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= tmp {
			left++
		}
		nums[right] = nums[left]
	}
	if left+1 == k {
		return tmp
	}
	nums[left] = tmp

	if left+1 > k {
		return find(start, left-1, nums, k)
	}
	return find(left+1, end, nums, k)
}

func TestFindKthLargest(t *testing.T) {
	s := []int{2, 6, 1, 8, 3, 5, 9, 10}
	k := 6

	t.Log(findKthLargest(s, k))
}
