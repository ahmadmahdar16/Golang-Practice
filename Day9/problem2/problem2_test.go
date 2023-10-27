package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinMax(t *testing.T) {
	a1 := 10
	a2 := 20
	a3 := 12
	a4 := 30
	a5 := 8
	a6 := 9
	t.Run("Case 1:", func(t *testing.T) {
		min, max := GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
		assert.Equal(t, 8, min, "Hasil output tidak sesuai")
		assert.Equal(t, 30, max, "Hasil output tidak sesuai")
	})

}
