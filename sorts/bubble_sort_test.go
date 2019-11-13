package sorts

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := []int{5, 3, 1, 6, 8, 10, 2}
	log.Println("in: ", input)
	BubbleSort(input)
	log.Println("out: ", input)
}
