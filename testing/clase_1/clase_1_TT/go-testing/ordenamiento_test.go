package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamientoGood(t *testing.T) {
	array := []int{8, 7, 9, 99, 46, 33}
	mustBeResult := []int{7, 8, 9, 33, 46, 99}
	result, err := Ordenamiento(array)
	assert.Equal(t, result, mustBeResult)
	assert.IsType(t, mustBeResult, result)
	assert.Nil(t, err)
}
func TestOrdenamientoBadParams(t *testing.T) {
	array := []int{}
	mustBeResult := []int{}
	result, err := Ordenamiento(array)
	assert.Equal(t, result, mustBeResult)
	assert.IsType(t, mustBeResult, result)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorSliceEmpty.Error())
}
