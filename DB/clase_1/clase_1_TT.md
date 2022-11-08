
# Bases de Datos Relacionales

## Diseño Diagrama Entidad-Relación

### Ejercicio

Se propone realizar las siguientes consultas a la base de datos movies_db.sql.

Importar el archivo movies_db.sql desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:

1. Mostrar todos los registros de la tabla de movies. 
    `
    SELECT movies.* FROM movies;
    `


2. Mostrar el nombre, apellido y rating de todos los actores.
    `
    SELECT a.first_name, a.last_name, a.rating FROM actors a;
    `

3. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español.
    `
    SELECT peliculas.title AS titulo FROM movies peliculas;
    `

4. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
    `
    SELECT a.first_name, a.last_name FROM actors a
    WHERE a.rating > 7.5;
    `

5. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios. 
    `
    SELECT m.title, rating FROM movies m
    WHERE m.rating > 7.5 AND m.awards > 2;
    `

6. Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
    `
    SELECT m.title, m.rating FROM movies m
    ORDER BY m.rating ASC;
    `

7. Mostrar los títulos de las primeras tres películas en la base de datos.
    `
    SELECT m.title FROM movies m
    LIMIT 3;
    `

8. Mostrar el top 5 de las películas con mayor rating.
    `
    SELECT movies.* FROM movies 
    ORDER BY movies.rating DESC
    LIMIT 5;
    `

9. Listar los primeros 10 actores.
    `
    SELECT actors.* FROM actors
    LIMIT 10;
    `

10. Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
    `
    SELECT movies.title, movies.rating FROM movies
    WHERE movies.title LIKE 'Toy Story%';
    `

11. Mostrar a todos los actores cuyos nombres empiezan con Sam.
    `
    SELECT actors.* FROM actors 
    WHERE actors.first_name LIKE 'Sam%';
    `

12. Mostrar el título de las películas que salieron entre el 2004 y 2008.
    `
    SELECT movies.title FROM movies 
    WHERE YEAR(movies.release_date) BETWEEN '2004' AND '2008';
    `

13. Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
    `
    SELECT m.title FROM movies m
    WHERE m.rating > 3 AND m.awards > 1 AND YEAR(m.release_date) BETWEEN '1988' AND '2009'
    ORDER BY m.rating DESC;
    `
 