/*
￼Ejercicio 3 - calculate salary
Una empresa marinera necesita calculate el salary de sus employees
basándose en la cantidad de horas trabajadas By mes y la categoría.

Si es categoría C, su salary es de $1.000 By hora
Si es categoría B, su salary es de $1.500 By hora más un %20 de su salary mensual
Si es de categoría A, su salary es de $3.000 By hora más un %50 de su salary mensual

Se solicita generar una función que reciba By parámetro la cantidad de
minutes trabajados By mes y la categoría,
y que devuelva su salary.
*/

package main

import "errors"
import "fmt"

var (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
)

func calculateSalaryByMinutes(minutes int, salary float64) float64 {
	return float64(minutes) * salary / 60
}

func salaryByEmployee(minutes int, category string) (float64, error) {
	var (
		salaryCategoryAByHora = 1000.00
		salaryCategoryBByHora = 1500.00 + 1500.00*20/100
		salaryCategoryCByHora = 3000.00 + 3000.00*50/100
	)
	if minutes < 0 {
		return 0, errors.New("minutes must be greater than 0")
	}
	switch category {
	case CategoryA:
		return calculateSalaryByMinutes(minutes, salaryCategoryAByHora), nil
	case CategoryB:
		return calculateSalaryByMinutes(minutes, salaryCategoryBByHora), nil
	case CategoryC:
		return calculateSalaryByMinutes(minutes, salaryCategoryCByHora), nil
	default:
		return 0, errors.New("category not found")
	}
}

func ejercicio3() {
	// Try the next lines
	// res, err := salaryByEmployee(9600, categoryB) // lun a vie 8hs = 9600 minutes
	// res, err := salaryByEmployee(3940, "D")
	res, err := salaryByEmployee(-8940, CategoryB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
