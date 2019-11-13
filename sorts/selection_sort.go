package sorts

func SelectionSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 0; i < len(a); i++ { // 遍历未排序区间
		// 查找最小值
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}
