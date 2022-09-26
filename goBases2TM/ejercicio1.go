/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, 
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y 
si gana más de $150.000 se le descontará además un 10%.
 */

package main

import "fmt"

func impuestoSalarial(salario float64)float64 {
	if(salario > 50000) {
		return salario*17/100
	} else {
		return salario*27/100
	}
}

 func main(){
	fmt.Println(impuestoSalarial(35000))
	fmt.Println(impuestoSalarial(200000.60))
 }

