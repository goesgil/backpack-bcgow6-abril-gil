// Cambiar el método para que no sólo retorne un entero sino también un error.
// Incorporar una validación en la que si el denominador es igual a 0,
// retorna un error cuyo mensaje sea 'El denominador no puede ser 0'.
// Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.

package calculadora

import (
	"fmt"
	"sort"
)

var (
	errorSliceEmpty = fmt.Errorf("error, slice empty")
)

func Ordenamiento(array []int) ([]int, error) {
	if len(array) == 0 {
		return []int{}, errorSliceEmpty
	}
	sort.Ints(array)
	return array, nil
}