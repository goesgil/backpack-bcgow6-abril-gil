/*
Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

package main

import (
	"fmt"
)

type NewError struct{}

func (e *NewError) NewError(salary int) error {
	return fmt.Errorf("error: el salario ingresado no alcanza el mínimo imponible, el mínimo imponible es de 150.000 y el salario ingresado es de: %v ", salary)
}

func ApplyImpuestBySalary(salary int) {
	var newErrorGenerated *NewError
	if salary < 150000 {
		fmt.Println(newErrorGenerated.NewError(salary))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func main() {
	salary := 30000
	ApplyImpuestBySalary(salary)
	salary = 230000
	ApplyImpuestBySalary(salary)
}
