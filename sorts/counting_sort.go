package sorts

import "math"

// CountingSort 计数排序
// 2, 5, 3, 0, 2, 3, 0, 3
func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 求最大值
	var max int = math.MinInt32
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	// 用另外一个数组c计数
	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}

	// 顺序求和, c[i]里面存储的就是小于等于i的元素的个数
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	// 遍历原数组, 计算出排序后的位置
	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}
