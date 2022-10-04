/* 
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros 
y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

*/

 package main
 import "errors"
 import "fmt"


func promedioPorAlumno(calificaciones ...int) (int, error) {
	var total int
	for _, calif := range calificaciones {
		if(calif < 0){
			return 0, errors.New("no se debe ingresar calificaciones menores a 0")
		}
		total+= calif
	}
	return total / len(calificaciones), nil
}

 func ejercicio2(){
	 res, err := promedioPorAlumno(5,7,3,7,7,8,9,-8,10,10)
	 if (err != nil) {
		 fmt.Println(err)
	 } else {
		 fmt.Println(res)
	 }
 }