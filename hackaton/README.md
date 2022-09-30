#  Hackathon Go Bases💥
Repositorio base para hackaton - Go Bases

## Objetivo
El objetivo de esta guía práctica es que podamos afianzar y profundizar los conceptos vistos en Go Bases. Para esto, vamos a plantear un desafío integrador que engloba los temas estudiados. 
Tendrán tiempo para realizarlo hasta las 10 hs AR/UR/CH - 8hs CO/ME. de mañana VIERNES

## Planteo:
Una aerolínea pequeña necesita un sistema de reservas de pasajes a diferentes países, y requiere un archivo con la información de los pasajes sacados en las últimas 24 horas. 
Deben crear un sistema de reservas para generar esos archivos
El archivo en cuestión es del tipo valores separados por coma (csv por su siglas en inglés), donde los campos están compuestos por: id, nombre, email, país de destino, hora del vuelo y precio. 

¿Are you ready? 



## Desafío
Realizar un programa que sirva como herramienta para crear, leer, actualizar y eliminar un ticket.  Para lograrlo, deben descargar el siguiente archivo .zip que será la base del proyecto.


¡Atención! Los ejemplos a continuación son sólo de guía, el desafío se puede resolver de múltiples maneras.



Requerimiento 1 - Cargar archivos: 
Implementar un módulo para leer el archivo donde se encuentran los tickets del día.


Requerimiento 2 - Crear: 
Un método para crear un nuevo ticket añadir al registro.


func (b *bookings) Create(t Ticket) (Ticket, error) {}


Requerimiento 3 - Leer: 
Un método para traer un ticket a través de su campo id.


func (b *bookings) Read(id int) (Ticket, error) {}


Requerimiento 4 - Update: 
Un método para actualizar los campos de un ticket.


func (b *bookings) Update(id int, t Ticket) (Ticket, error) {}


Requerimiento 5 - Delete: 
Un método para eliminar un campo por su id.


func (b *bookings) Delete(id int) (int, error) {}
