/*

Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

	1	perro necesitan 10 kg de alimento
	2	gato 5 kg
	3	Hamster 250 gramos.
	4	Tarántula 150 gramos.

Se solicita:
	1	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
	2	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/

package main

import (
	"errors"
	"fmt"
)

var (
	DogAnimal     = "DOG"
	CatAnimal     = "CAT"
	HamsterAnimal = "HAMSTER"
	SpiderAnimal  = "SPIDER"
)

func calculateFoodBySpider(quantityAnimal float64) (quantityFood float64) {
	return 0.150 * quantityAnimal
}

func calculateFoodByDog(quantityAnimal float64) (quantityFood float64) {
	return 10000 * quantityAnimal
}

func calculateFoodByCat(quantityAnimal float64) (quantityFood float64) {
	return 5000 * quantityAnimal
}

func calculateFoodByHamster(quantityAnimal float64) (quantityFood float64) {
	return 0.250 * quantityAnimal
}

func FindAlimentByAnimal(typeAnimal string, quantityAnimal int) (quantityFood float64, err error) {
	switch typeAnimal {
	case SpiderAnimal:
		return calculateFoodBySpider(float64(quantityAnimal)), nil
	case DogAnimal:
		return calculateFoodByDog(float64(quantityAnimal)), nil
	case CatAnimal:
		return calculateFoodByCat(float64(quantityAnimal)), nil
	case HamsterAnimal:
		return calculateFoodByHamster(float64(quantityAnimal)), nil
	default:
		return 0, errors.New("this animal not found")
	}
}
func main() {
	res, err := FindAlimentByAnimal(SpiderAnimal, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("El animal %s, necesita %.3fgr de alimento\n", SpiderAnimal, res)
}
