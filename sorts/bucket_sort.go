package sorts

import (
	"fmt"
)

// 获取待排序数组中的最大值
func getMax(a []int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// BucketSort 桶排序算法
// arr = [1, 6, 3, 5, 8, 6, 4] max:8 bucketsLength:7 index:0(取值范围[0, 6])
// index: arr[i] * (7-1) / 8
// buckets = {{1}, {}, {3}, {5, 4}, {6, 6}, {}, {8}}
// for i in buckets:
// 		桶内快排
// 		根据偏移量写回arr
//		更新arr偏移量
func BucketSort(a []int) {
	num := len(a)
	if num <= 1 {
		return
	}
	max := getMax(a)
	buckets := make([][]int, num) // 二维切片

	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max                // 桶序号
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])        // 桶内做快速排序
			copy(a[tmpPos:], buckets[i]) // 写回
			tmpPos += bucketLen          // 更新偏移量
		}
	}

}

// 桶排序简单实现
func BucketSortSimple(source []int) {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]]++
	}
	fmt.Println(array)
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)

}
