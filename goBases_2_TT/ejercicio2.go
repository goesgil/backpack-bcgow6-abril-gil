/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad
para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz,
la dimensión del alto, la dimensión del ancho, si es cuadrática
y cuál es el valor máximo.

*/

package main

import "fmt"

type Matrix struct {
	Height      float64
	Width       float64
	ValueMax    float64
	IsCuadratic bool
}

func (m *Matrix) Set(height, width, valueMax float64, isCuadratic bool) {
	newMatrix := m
	newMatrix.Height = height
	newMatrix.Width = width
	newMatrix.ValueMax = valueMax
	newMatrix.IsCuadratic = isCuadratic
}

func (m Matrix) Print() {
	fmt.Printf("Altura: %v\n", m.Height)
	fmt.Printf("Tamaño: %v\n", m.Width)
	fmt.Printf("Valor maximo: %v\n", m.ValueMax)
	fmt.Printf("Es cuadratica: %v\n", m.IsCuadratic)
}

func main() {
	var newMatric Matrix
	newMatric.Set(
		13.5,
		23,
		400,
		false,
	)
	newMatric.Print()
}
