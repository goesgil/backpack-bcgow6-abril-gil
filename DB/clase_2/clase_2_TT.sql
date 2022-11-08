-- Ejercicio 1
-- Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y 
-- guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
-- Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.

CREATE TEMPORARY TABLE TWD
SELECT e.* FROM episodes e
INNER JOIN seasons
ON seasons.id = e.season_id
WHERE seasons.serie_id IN (SELECT s.id FROM series s WHERE s.title LIKE "The Walking Dead%");

SELECT TWD.* FROM TWD;

-- Ejercicio 2
-- En la base de datos “movies”, seleccionar una tabla donde crear un índice 
-- y luego chequear la creación del mismo. 

ALTER TABLE actors
ADD INDEX idx_last_name (last_name);

SHOW INDEX FROM actors;
   
