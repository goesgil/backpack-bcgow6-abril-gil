/*

Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ejercicio2() {
	res, err := os.ReadFile("./listProduct.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	values := strings.Split(string(res), "\n")
	priceTotal := 0.00
	for i := 0; i < len(values)-1; i++ {
		valuesIndividual := strings.Split(values[i], ";")
		if priceToInt, err := strconv.ParseFloat(valuesIndividual[1], 64); i!=0 && err == nil {
			priceTotal += priceToInt
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		
		fmt.Printf("%s \t %s \t %s\n", valuesIndividual[0], valuesIndividual[1], valuesIndividual[2])
	}
	fmt.Printf("\t %.2f \t\n", priceTotal)
}
