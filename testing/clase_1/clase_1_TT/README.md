
# **Programación Go** CLASE 1


# **Testing**

**Práctica clase 1 - Go Testing**

## **Objetivo**
Implementar correctamente las librerías testing y testify para la creación de test unitarios. Efectuar validaciones correctamente sobre funciones y métodos específicos. Y así generar un primer acercamiento práctico con test unitario en golang.

## **Forma de trabajo**
Deben diseñar test unitarios según lo planteado en los ejercicios. Cada test debe implementar la librería testify. Los mismos deben ser realizados en sus computadoras. Deben generar una carpeta llamada go-testing. 

___

### Ejercicio 1 - _Test Unitario Restar_

Para el método Restar() visto en la clase, realizar el test unitario correspondiente.
Para esto:

Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.

Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.

___

### Ejercicio 2 - _Test Unitario Método Ordenar_

Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente diseñar un test unitario que valide el funcionamiento del mismo.

Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.

Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.


___

### Ejercicio 3 - _Test Unitario Método Dividir_
Para el Método Dividir, visto en la clase:

Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retorna un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.

Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.

Dentro de la carpeta go-testing crear un archivo dividir test.go con el test diseñado.
