package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	a := 10
	b := 20
	t.Run("Case 1:", func(t *testing.T) {
		x, y := Swap(&a, &b)
		assert.Equal(t, 20, x, "Hasil output tidak sesuai")
		assert.Equal(t, 10, y, "Hasil output tidak sesuai")
	})

}
