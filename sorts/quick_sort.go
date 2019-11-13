package sorts

// QuickSort 快速排序: 分治思想, 递归实现, 不稳定排序, 时间复杂度O(nlogn), 最坏时间复杂度O(n^2), 原地排序
// 递推公式: quick_sort(p...r) = quick_sort(p...q-1) + quick_sort(q+1...r)
// 终止条件: p >= r
func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

// arr = [11, 8, 3, 9, 7, 1, 2, 5] pivot:5 i:0 j:0 start:0 end:7
// j:0 [11, 8, 3, 9, 7, 1, 2, 5] </> i:0
// j:1 [11, 8, 3, 9, 7, 1, 2, 5] </> i:0
// j:2 [3, 8, 11, 9, 7, 1, 2, 5] arr[2]<->arr[0] i:1
// j:3 [3, 8, 11, 9, 7, 1, 2, 5] </> i:1
// j:4 [3, 8, 11, 9, 7, 1, 2, 5] </> i:1
// j:5 [3, 1, 11, 9, 7, 8, 2, 5] arr[5]<->arr[1] i:2
// j:6 [3, 1, 2, 9, 7, 8, 11, 5] arr[6]<->arr[2] i:3
// break [3, 1, 2, 5, 7, 8, 11, 9] arr[i]<->arr[end]
// i:3 => separateSort(arr, 0, 2), separateSort(arr, 4, 7)
// separateSort(arr, 0, 2) => i:1 => separateSort(arr, 0, 0), separateSort(arr, 2, 2) => return, return => [1, 2, 3, 5, 7, 8, 11, 9]
// separateSort(arr, 4, 7) => i:6 => separateSort(arr, 4, 5), separateSort(arr, 7, 7) => [1, 2, 3, 5, 7, 8, 9, 11]
// separateSort(arr, 4, 5) => i:4 => separateSort(arr, 4, 3), separateSort(arr, 5, 5) => return, return => [1, 2, 3, 5, 7, 8, 9, 11]
// separateSort(arr, 7, 7) => return => [1, 2, 3, 5, 7, 8, 9, 11]
// END
func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end) // 获取分区点
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
