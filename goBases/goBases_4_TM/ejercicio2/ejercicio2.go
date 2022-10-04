/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"errors"
	"fmt"
)

type NewError struct{}

func (e *NewError) NewError() error {
	return errors.New("error: el salario ingresado no alcanza el mínimo imponible")

}

func ApplyImpuestBySalary(salary int) (string, error) {
	var newErrorGenerated *NewError
	if salary < 150000 {
		return "", newErrorGenerated.NewError()
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 30000
	res, err := ApplyImpuestBySalary(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	salary = 230000
	fmt.Println(ApplyImpuestBySalary(salary))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
