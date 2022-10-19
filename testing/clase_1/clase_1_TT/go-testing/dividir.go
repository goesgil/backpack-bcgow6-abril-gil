// Cambiar el método para que no sólo retorne un entero sino también un error.
// Incorporar una validación en la que si el denominador es igual a 0,
// retorna un error cuyo mensaje sea 'El denominador no puede ser 0'.
// Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.

package calculadora

import "fmt"

var (
	errorDivisor = fmt.Errorf("error, divisor must be greater than zero")
)

func Dividir(a, b int) (int, error) {
	if b < 1 {
		return 0, errorDivisor
	}
	return a / b, nil
}
