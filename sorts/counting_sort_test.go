package sorts

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{5, 4}
	CountingSort(arr, len(arr))
	t.Log(arr)

	arr = []int{2, 5, 3, 0, 2, 3, 0, 3}
	CountingSort(arr, len(arr))
	t.Log(arr)
}
