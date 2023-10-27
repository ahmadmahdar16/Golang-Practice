package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudentsScore(t *testing.T) {
	var a = Student{}
	a.name = append(a.name, "andi")
	a.score = append(a.score, 90)
	a.name = append(a.name, "budi")
	a.score = append(a.score, 80)
	a.name = append(a.name, "cindy")
	a.score = append(a.score, 100)
	a.name = append(a.name, "dodi")
	a.score = append(a.score, 70)
	a.name = append(a.name, "endi")
	a.score = append(a.score, 60)
	a.name = append(a.name, "fikri")
	a.score = append(a.score, 90)
	t.Run("Case 1:", func(t *testing.T) {
		avg := a.Average()
		scoreMax, nameMax := a.Max()
		scoreMin, nameMin := a.Min()
		assert.Equal(t, float64(81.66666666666667), avg, "Hasil output tidak sesuai")
		assert.Equal(t, 100, scoreMax, "Hasil output tidak sesuai")
		assert.Equal(t, "cindy", nameMax, "Hasil output tidak sesuai")
		assert.Equal(t, 60, scoreMin, "Hasil output tidak sesuai")
		assert.Equal(t, "endi", nameMin, "Hasil output tidak sesuai")
	})

}
