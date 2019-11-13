package sorts

func BubbleSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 0; i < len(a); i++ {
		// 提前退出标志
		flag := false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				// 有效数据交换
				flag = true
			}
		}
		// 如果没有数据交换, 提前退出
		if !flag {
			break
		}
	}
}
