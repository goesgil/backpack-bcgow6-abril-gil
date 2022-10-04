/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad
para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha
y que tenga un método detalle

*/

package main

import (
	"fmt"
	"strconv"
)

type alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a alumno) Detalle() {
	fmt.Println("Nombre: " + a.Nombre)
	fmt.Println("Apellido: " + a.Apellido)
	fmt.Println("DNI: " + strconv.Itoa(a.DNI))
	fmt.Println("Fecha: " + a.Fecha)
}
func ejercicio1() {
	alumno1 := alumno{
		Nombre:   "Leo",
		Apellido: "Dicaprio",
		DNI:      34653789,
		Fecha:    "11/03/1993",
	}
	alumno1.Detalle()
}
