package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyChange(t *testing.T) {
	type TestCase struct {
		Param    int
		Expected []int
		Message  string
	}

	arrTestCase := []TestCase{
		{123, []int{100, 20, 1, 1, 1}, "Hasil Tidak Sesuai"},
		{432, []int{200, 200, 20, 10, 1, 1}, "Hasil Tidak Sesuai"},
		{543, []int{500, 20, 20, 1, 1, 1}, "Hasil Tidak Sesuai"},
		{7752, []int{5000, 2000, 500, 200, 50, 1, 1}, "Hasil Tidak Sesuai"},
		{15321, []int{10000, 5000, 200, 100, 20, 1}, "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, MoneyChange(arrTestCase[i].Param), arrTestCase[i].Message)
		})
	}

}
