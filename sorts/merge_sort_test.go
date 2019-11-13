package sorts

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{11, 8, 3, 9, 7, 1, 2, 5}
	log.Println("in: ", input)
	MergeSort(input)
	log.Println("out: ", input)
}
