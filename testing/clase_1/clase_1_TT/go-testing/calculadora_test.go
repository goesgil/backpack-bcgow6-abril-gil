package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestarGood(t *testing.T) {
	// Assert that
	numA := 44
	numB := 4
	resultMock := 40

	// Act
	result := Restar(numA, numB)

	// Add
	assert.Equal(t, result, resultMock)
}
