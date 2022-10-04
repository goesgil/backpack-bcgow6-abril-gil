/*

Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

  var employees = map[string]int{
	"Benjamin": 20,
	"Nahuel": 26,
	"Brenda": 19,
	"Darío": 44,
	"Pedro": 30
}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.

*/

package main

import "fmt"

func ejercicio4() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30}
	nameEmployeeSearch := "Pedro"
	fmt.Printf("Busqueda de empleado: %v , tiene %v años\n",
		nameEmployeeSearch,
		employees[nameEmployeeSearch])

	count := 0
	for _, v := range employees {
		if v > 21 {
			count++
		}
	}
	fmt.Printf("Cantidad de empleados mayor de 21 años: %v\n",
		count)

	employees["Federico"] = 25
	fmt.Printf("Nuevo empleado agregado a la lista\n")

	delete(employees,"Pedro")
	fmt.Printf("Empleado Pedro eliminado. Lista actualizada %v \n", employees)
}
