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

func main() {

	productOne := products{
		id:       20,
		price:    26.7,
		quantity: 19,
	}
	productTwo := products{
		id:       20,
		price:    26.5,
		quantity: 19,
	}

	var newListProducts listProducts

	newListProducts.list = append(newListProducts.list, productOne, productTwo)

	var listProductToString string

	for _, i := range newListProducts.list {
		listProductToString += fmt.Sprintf("%d;%d;%.2f;", i.id, i.quantity, i.price)
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
