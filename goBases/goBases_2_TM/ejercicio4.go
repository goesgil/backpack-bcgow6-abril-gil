/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular
algunas estadísticas de calificaciones de los alumnos de un curso,
requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
(mínimo, máximo o promedio)
y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido)
que se le puede pasar una cantidad N de enteros
y devuelva el cálculo que se indicó en la función anterior
*/

package main

import (
	"errors"
	"fmt"
)

var (
	CalculateMax     = "CALCULATE_MAX"
	CalculateMin     = "CALCULATE_MIN"
	CalculatePromedy = "CALCULATE_PROMEDY"
)

func calculatePromedyMaxCalifications(califications []int) (maxCalification int) {
	calificationMax := califications[0]
	for _, c := range califications {
		if c > calificationMax {
			calificationMax = c
		}
	}
	return calificationMax
}

func calculatePromedyMinCalifications(califications []int) (minCalifications int) {
	calificationMin := califications[0]
	for _, c := range califications {
		if c < calificationMin {
			calificationMin = c
		}
	}
	return calificationMin
}

func calculatePromedyCalifications(califications []int) (promedyCalifications int) {
	calculationPromedy := 0
	for _, c := range califications {
		calculationPromedy += c
	}
	return calculationPromedy/len(califications)
}

func CalculateCalificationsByType(califications []int, typeCalculate string) (resultPromedy int, err error) {
	if len(califications) == 0 {
		return 0, errors.New("quantity califications must be greater than 0")
	}
	switch typeCalculate {
	case CalculateMax:
		return calculatePromedyMaxCalifications(califications), nil
	case CalculateMin:
		return calculatePromedyMinCalifications(califications), nil
	case CalculatePromedy:
		return calculatePromedyCalifications(califications), nil
	default:
		return 0, errors.New("type of calculate not found")
	}
}

func ejercicio4() {
	calificationsStudent1 := []int{4, 6, 7, 9, 8, 9, 10, 9}
	res, err := CalculateCalificationsByType(calificationsStudent1, CalculatePromedy)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
