## Laboratorio 3 - Sistemas Distribuidos

#### Leia Organa:
- Consulta al servidor Broker la cantidad de rebeldes que hay en tal ciudad. Recibe por respuesta la cantidad de rebeldes que haya según el fulcrum consultado. 
- Monotonic Reads se ve reflejada en la comparación de vectores nuevos con los últimos consultados, se ve cuál es más nuevo según el servidor del que venga, si es más actualizado, se devuelve el número de rebeldes que llegó en el request, si es menor (por lo tanto menos actualizado) se imprime el valor antes consultado.

