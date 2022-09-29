/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/

package main

import "fmt"

type product struct {
	name     string
	price    float64
	quantity int
}

func (p *product) CreateProduct(name string, price float64) product {
	p.name = name
	p.price = price
	return *p
}

type Product interface {
	CreateProduct(name string, price float64)
}

type User interface {
	AddProducts()
	DeleteProducts()
}

type user struct {
	GivenName  string
	FamilyName string
	Email      string
	Products   []product
}

func (u *user) AddProducts(productToAdd *product, quantity int) {
	productToAdd.quantity = quantity
	u.Products = append(u.Products, *productToAdd)
}

func (u *user) DeleteProducts(newUser user) {
	u.Products = []product{}
}

func AddProducts(u *user, p *product, quantity int) {
	u.AddProducts(p, quantity)
}

func main() {
	newUser := user{
		GivenName:  "Abril",
		FamilyName: "Gil",
	}
	productOne := product{
		name:  "Pava",
		price: 500.7,
	}
	fmt.Println(newUser)
	AddProducts(&newUser, &productOne, 4)
	fmt.Println((newUser))

}
