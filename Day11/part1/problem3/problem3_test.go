package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDragonOfLoowater(t *testing.T) {
	type TestCase struct {
		Param    [][]int
		Expected string
		Message  string
	}

	arrTestCase := []TestCase{
		{[][]int{[]int{5, 4}, []int{7, 8, 4}}, "11", "Hasil Tidak Sesuai"},
		{[][]int{[]int{5, 10}, []int{5}}, "Knight Fall", "Hasil Tidak Sesuai"},
		{[][]int{[]int{7, 2}, []int{4, 3, 1, 2}}, "Knight Fall", "Hasil Tidak Sesuai"},
		{[][]int{[]int{7, 2}, []int{2, 1, 8, 5}}, "10", "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, DragonOfLoowater(arrTestCase[i].Param[0], arrTestCase[i].Param[1]), arrTestCase[i].Message)
		})
	}

}
