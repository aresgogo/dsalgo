package sorts

import (
	"log"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	input := []int{5, 3, 1, 6, 8, 10}
	log.Println("in: ", input)
	SelectionSort(input)
	log.Println("out: ", input)
}
