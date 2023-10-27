package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstitutionChiper(t *testing.T) {
	var a = student{}
	var c Chiper = &a
	a.name = "fakhry"
	a.nameEncode = "idnkub"
	t.Run("Case 1:", func(t *testing.T) {
		assert.Equal(t, "idnkub", c.Encode(), "Hasil output tidak sesuai")
		assert.Equal(t, "fakhry", c.Decode(), "Hasil output tidak sesuai")
	})

}
