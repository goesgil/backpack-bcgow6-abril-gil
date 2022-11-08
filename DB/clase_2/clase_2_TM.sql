#Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT empleado.nombre, empleado.puesto, departamento.localidad 
FROM empleado, departamento
WHERE empleado.depto_nro = departamento.depto_nro;

#Visualizar los departamentos con más de cinco empleados.
SELECT d.depto_nro, d.nombre_depto, d.localidad FROM departamento d
INNER JOIN empleado e
ON e.depto_nro = d.depto_nro
GROUP BY d.depto_nro
HAVING COUNT(d.depto_nro) > 2;

#Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
SELECT e.nombre, e.salario,e.depto_nro FROM empleado e
INNER JOIN departamento d
ON d.depto_nro = e.depto_nro
WHERE e.puesto = "Presidente";

#Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
SELECT e.nombre, e.apellido, e.puesto, e.salario, e.comision, e.fecha_alta FROM empleado e
INNER JOIN departamento d
ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto = "Contabilidad"
ORDER BY e.nombre ASC;

#Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT empleado.nombre FROM empleado 
WHERE empleado.salario IN (SELECT MIN(salario) FROM empleado);

#Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
-- Corregir este ejercicio ---------------------------------------------------------------
SELECT e.nombre, e.salario FROM empleado e
INNER JOIN departamento d
ON d.depto_nro = e.depto_nro
WHERE d.nombre_depto LIKE "Ventas" AND e.salario IN (SELECT MIN(salario) FROM empleado);
-- ---------------------------------------------------------------------------------------

