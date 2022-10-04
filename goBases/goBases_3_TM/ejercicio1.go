/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/

package main

import (
	"fmt"
	"os"
)

type products struct {
	id       int
	price    float64
	quantity int
}
type listProducts struct {
	list []products
}

func ejercicio1() {

	productOne := products{
		id:       20,
		price:    500.7,
		quantity: 19,
	}
	productTwo := products{
		id:       20,
		price:    1200.5,
		quantity: 19,
	}
	productThree := products{
		id:       23440,
		price:    1000,
		quantity: 30,
	}

	var newListProducts listProducts

	newListProducts.list = append(newListProducts.list, productOne, productTwo, productThree)

	var listProductToString = "ID;PRICE;QUANTITY;\n"

	for _, i := range newListProducts.list {
		listProductToString += fmt.Sprintf("%d;%.2f;%d;\n", i.id, i.price, i.quantity)
	}

	b := []byte(listProductToString)

	err := os.WriteFile("./listProduct.csv", b, 777)

	if err != nil {
		fmt.Println(err)
	}

	res, err := os.ReadFile("./listProduct.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
