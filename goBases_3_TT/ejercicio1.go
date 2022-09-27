/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.

*/

package main

import "fmt"

type user struct {
	GivenName  string
	FamilyName string
	YearsOld   int
	Email      string
	password   string
}

func (users *user) ChangeName(givenName, familyName string) {
	users.GivenName = givenName
	users.FamilyName = familyName
}
func (users *user) ChangeYearsOld(yearsOld int) {
	users.YearsOld = yearsOld
}
func (users *user) ChangeEmail(email string) {
	users.Email = email
}
func (users *user) ChangePassword(newPassword string) {
	users.password = newPassword
}

func main() {
     newUser := user{
        GivenName: "Abril",
        FamilyName: "Gil",
    }
    oldUser := &newUser

	fmt.Println(newUser)
	fmt.Println(oldUser)
    oldUser.ChangeName("Pablo","Saramillo")

	fmt.Println(newUser)
	fmt.Println(oldUser)
    
}
