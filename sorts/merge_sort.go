package sorts

// MergeSort 归并排序: 分治思想, 递归实现, 稳定排序, 时间复杂度O(nlogn), 空间复杂度O(n^2)
// 递推公式: merge_sort(p...r) = merge(merge_sort(p...q), merge_sort(q+1...r))
// 终止条件: p >= r 不用再继续分解
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

// arr = [11, 8, 3, 9, 7, 1, 2, 5]
// mergeSort(arr, 0, 7)
// mergeSort(arr, 0, 3) mergeSort(4, 7)
// mergeSort(arr, 0, 1) mergeSort(arr, 2, 3) mergeSort(arr, 4, 5) mergeSort(arr, 6, 7)
// mergeSort(arr, 0, 0) mergeSort(arr, 1, 1) mergeSort(arr, 2, 2) mergeSort(3, 3) mergeSort(4, 4) mergeSort(5, 5) mergeSort(6, 6) merge(7, 7)
// return 				return 				 return 			  return 		  return 		  return 		  return 		  return
// merge(arr, 0, 0, 1) merge(arr, 2, 2, 3) merge(arr, 4, 4, 5) merge(arr, 6, 6, 7)
// merge(arr, 0, 1, 3) merge(4, 5, 7)
// merge(arr, 0, 3, 7)
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}
