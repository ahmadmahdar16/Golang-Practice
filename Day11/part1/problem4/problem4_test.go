package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	type TestCase struct {
		Param1   []int
		Param2   int
		Expected int
		Message  string
	}

	arrTestCase := []TestCase{
		{[]int{1, 1, 3, 5, 5, 6, 7}, 3, 2, "Hasil Tidak Sesuai"},
		{[]int{1, 2, 3, 5, 6, 8, 10}, 5, 3, "Hasil Tidak Sesuai"},
		{[]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53, 6, "Hasil Tidak Sesuai"},
		{[]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100, -1, "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, BinarySearch(arrTestCase[i].Param1, arrTestCase[i].Param2), arrTestCase[i].Message)
		})
	}
}
