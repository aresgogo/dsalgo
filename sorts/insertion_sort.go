package sorts

func InsertionSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 1; i < len(a); i++ { // 遍历未排序区间
		value := a[i]       // 获取未排序区间起始位置的值
		j := i - 1          // 获取已排序区间结束位置的下标
		for ; j >= 0; j-- { // 遍历已排序区间
			if a[j] > value { // 大于要插入的值
				a[j+1] = a[j] // 数据移动
			} else {
				break
			}
		}
		a[j+1] = value // 插入数据
	}
}
