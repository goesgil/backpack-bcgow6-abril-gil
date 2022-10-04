/* Ejercicio 3 - Declaración de variables

Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.

  var 1nombre string
  var apellido string
  var int edad
  1apellido := 6
  var licencia_de_conducir = true
  var estatura de la persona int
  cantidadDeHijos := 2

*/

package main

import "fmt"

func ejercicio3() {
	var nombre string
	var edad int
	var apellido = "6"
	var licenciaDeConducir = true
	var estaturaDeLaPersona int

	fmt.Println(
		nombre,
		apellido,
		edad,
		licenciaDeConducir,
		estaturaDeLaPersona,
	)
}
