package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirGood(t *testing.T) {
	numA := 100
	numB := 4
	mustBeResult := 25
	result, err := Dividir(numA, numB)
	assert.Equal(t, result, mustBeResult)
	assert.IsType(t, mustBeResult, result)
	assert.Nil(t, err)
}
func TestDividirBadParams(t *testing.T) {
	numA := 100
	numB := -4
	mustBeResult := 0
	result, err := Dividir(numA, numB)
	assert.Equal(t, result, mustBeResult)
	assert.IsType(t, mustBeResult, result)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorDivisor.Error())
}