package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffSlice(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	b := []string{"b", "d"}

	diff := DiffSlice(a, b)

	assert.Equal(t, diff, []string{"a", "c"}, "DiffSlice error")
}
