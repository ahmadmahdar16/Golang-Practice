package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleEquations(t *testing.T) {
	type TestCase struct {
		Param    []int
		Expected interface{}
		Message  string
	}

	arrTestCase := []TestCase{
		{[]int{1, 2, 3}, "no solution", "Hasil Tidak Sesuai"},
		{[]int{6, 6, 14}, []int{1, 2, 3}, "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		paramRune := arrTestCase[i].Param
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, SimpleEquations(paramRune[0], paramRune[1], paramRune[2]), arrTestCase[i].Message)
		})
	}

}
