package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntFields(t *testing.T) {
	s := "1 2 3 4"
	ns, err := IntFields(&s)
	assert.Nil(t, err)
	assert.EqualValues(t, ns, []int{1, 2, 3, 4})
}

func TestStringsToInts(t *testing.T) {
	ns, err := StringsToInts([]string{"1", "2", "3", "4"})
	assert.Nil(t, err)
	assert.EqualValues(t, ns, []int{1, 2, 3, 4})
}
