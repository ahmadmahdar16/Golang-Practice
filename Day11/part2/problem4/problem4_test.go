package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	type TestCase struct {
		Param    int
		Expected string
		Message  string
	}

	arrTestCase := []TestCase{
		{6, "VI", "Hasil Tidak Sesuai"},
		{9, "IX", "Hasil Tidak Sesuai"},
		{23, "XXIII", "Hasil Tidak Sesuai"},
		{2021, "MMXXI", "Hasil Tidak Sesuai"},
		{1646, "MDCXLVI", "Hasil Tidak Sesuai"},
	}

	for i := 0; i < len(arrTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, arrTestCase[i].Expected, RomanNumerals(arrTestCase[i].Param), arrTestCase[i].Message)
		})
	}
}
