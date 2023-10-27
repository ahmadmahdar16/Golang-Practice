package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrog(t *testing.T) {
	type TestCase struct {
		Param    []int
		Expected int
		Message  string
	}

	arrTestCase := []TestCase{
		{[]int{10, 30, 40, 20}, 30, "Hasil Tidak Sesuai"},
		{[]int{30, 10, 60, 10, 60, 50}, 40, "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, Frog(arrTestCase[i].Param), arrTestCase[i].Message)
		})
	}
}
