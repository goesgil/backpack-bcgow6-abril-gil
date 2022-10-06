# **GO WEB CLASE 4 TM**

## **Ejercicio 1 - Manejo de respuestas genéricas**


Se requiere implementar un manejo de respuestas genéricas para enviar siempre el mismo formato en las peticiones. Para lograrlo se deben realizar los siguientes pasos:
1. Generar el paquete web dentro del directorio pkg.
2. Realizar la estructura Response con los capos: code, data y error:
    
    a. code tendra el codigo de retorno.
    
    b. data tendrá la estructura que envía la aplicación (en caso que no haya error).
    
    c. error tendrá el error recibido en formato texto (en caso que haya error).

3. Desarrollar una función que reciba el code cómo entero, data como interfaz y error como string.
4. La función debe retornar en base al código, si es una respuesta con el data o con el error.
5. Implementar esta función en todos los retornos de los controladores, antes de enviar la respuesta al cliente la función debe generar la estructura que definimos.

