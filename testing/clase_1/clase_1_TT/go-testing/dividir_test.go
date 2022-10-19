// Cambiar el método para que no sólo retorne un entero sino también un error.
// Incorporar una validación en la que si el denominador es igual a 0,
// retorna un error cuyo mensaje sea 'El denominador no puede ser 0'.
// Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.

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