package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopDown(t *testing.T) {
	type TestCase struct {
		Param    int
		Expected int
		Message  string
	}

	arrTestCase := []TestCase{
		{0, 0, "Hasil Tidak Sesuai"},
		{1, 1, "Hasil Tidak Sesuai"},
		{2, 1, "Hasil Tidak Sesuai"},
		{3, 2, "Hasil Tidak Sesuai"},
		{5, 5, "Hasil Tidak Sesuai"},
		{6, 8, "Hasil Tidak Sesuai"},
		{7, 13, "Hasil Tidak Sesuai"},
		{9, 34, "Hasil Tidak Sesuai"},
		{10, 55, "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, TopDown(arrTestCase[i].Param), arrTestCase[i].Message)
		})
	}
}
