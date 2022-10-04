/*
Ejercicio 2 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos para registrar a un cliente son:
Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos. Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”. Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y abortc.
Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existc. Para ello, necesitas leer los datos de un archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error). Debes manipular adecuadamente ese error como hemos visto hasta aquí.
Ese error deberá:
1.-   generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmentc.
Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden). Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la validación pertinente para el caso de error retornado).
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
)

type RequestCreateClient struct {
	GivenName  string
	FamilyName string
	DNI        int
	Telephone  int
	Address    string
}

type client struct {
	File       int
	GivenName  string
	FamilyName string
	DNI        int
	Telephone  int
	Address    string
}

func (c *client) CreateClient(id int, req *RequestCreateClient) {
	c.File = id
	c.FamilyName = req.FamilyName
	c.GivenName = req.GivenName
	c.Address = req.Address
	c.DNI = req.DNI
	c.Telephone = req.Telephone
}

type Client interface {
	CreateClient(id int, req *RequestCreateClient)
}

func NewClient() Client {
	return &client{}
}

type CustomError struct{}

func (e *CustomError) WasNotGeneratedId() error {
	return fmt.Errorf("error generating ID")
}

func (e *CustomError) FileNotExist() error {
	return fmt.Errorf("el archivo indicado no fue encontrado o está dañado")
}

func (e *CustomError) InvalidRequest() error {
	return fmt.Errorf("Invalid request")
}

func GenerateId() (int, error) {
	err := CustomError{}
	id := rand.Int()
	if id == 0 {
		return id, err.WasNotGeneratedId()
	} else {
		return id, nil
	}
}

func ReaderFileCustomers(pathFile string) (string, error) {
	newError := CustomError{}
	file, err := os.ReadFile(pathFile)
	if err != nil {
		return "", newError.FileNotExist()
	}
	return string(file), nil
}

func ValidateRequest(
	c *RequestCreateClient) error {
	err := CustomError{}
	if c.DNI < 1 || c.Telephone < 1 || len(c.GivenName) < 1 || len(c.FamilyName) < 1 || len(c.Address) < 1 {
		return err.InvalidRequest()
	}
	return nil
}

func ejercicio2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("Fin de la ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()
	pathFile := "./customers.txt"
	requestToCreateClient := &RequestCreateClient{
		DNI:        23000111,
		Telephone:  2234453000,
		GivenName:  "Maar",
		FamilyName: "GoesGil",
		Address:    "Av. Siempre Viva 123",
	}
	if err := ValidateRequest(requestToCreateClient); err != nil {
		panic(err)
	}
	newClients := NewClient()

	ReaderFileCustomers(pathFile)

	id, err := GenerateId()
	if err != nil {
		panic(err)
	}
	newClients.CreateClient(id, requestToCreateClient)

	fmt.Println(newClients)
}
