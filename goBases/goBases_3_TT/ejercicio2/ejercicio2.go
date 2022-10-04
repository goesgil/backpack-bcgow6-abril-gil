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

func newProduc() product {
	return product{}
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

func (u *user) CreateUser(givenName, familyName string) {
	u.GivenName = givenName
	u.FamilyName = familyName
}


func newUser() user {
	return user{}
}

type Product interface {
	CreateProduct(name string, price float64)
}

type User interface {
	AddProducts()
	DeleteProducts()
}

func ejercicio2() {
	user1 := newUser()
	user1.CreateUser("Abril", "Gil")
	productOne := newProduc()
	productOne.CreateProduct("Laptop", 1000)
	fmt.Println(user1)
	user1.AddProducts(&productOne, 4)
	fmt.Println((user1))
}
