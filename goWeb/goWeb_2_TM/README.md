# **GO WEB** CLASE 2

## **Ejercicio 1 - Crear Entidad**
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).

## **Ejercicio 2 - Validación de campos**
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).

##  Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.