## Laboratorio 3 - Sistemas Distribuidos

#### Integrantes:

- Fernanda Cerda 201804567-5 Paralelo 200
- Alfredo Llanos 201804536-5 Paralelo 200
- Sofia Mañana 201804535-7 Paralelo 201


#### Distribución de máquinas virtuales y sus claves: 
| Máquina| Asignados| Claves|
| ----- | ---- | ---- |
| Dist 29| Fulcrum 1, Ahsoka Tano | hZK%8XqG-?Naj!>~|
| Dist 30| Fulcrum 2, Almirante Thrawn | M*jvU4W;$#DZs_5,|
| Dist 31| Fulcrum 3, Leia Organa | A3h!)7czKgsu{?GC|
| Dist 32| Broker Mos Eisley| a7+VT<F^*8mLvjX]|


#### Instrucciones para compilar:

- Debe abrir **7 terminales distintas**. En todas ellas debe dirigirse a la carpeta llamada "Lab3_SD". A continuación se explica qué abrir y cómo en cada una de ellas.
- En las dos primeras, conéctese a la máquina **"dist29"**, mediante la clave mencionada en la tabla de arriba. En una de ellas, dirígase a la carpeta **"Fulcrum"** y ejecute el archivo fulcrum.go mediante `go run fulcrum.go`. En la otra, dirígase a la carpeta **"Ahsoka"** y ejecute el archivo Ahsoka Tano.go mediante `go run Ahsoka Tano.go`.
- En las dos siguientes, conéctese a la máquina **"dist30"**. En una de ellas, dirígase a la carpeta **"Fulcrum2"** y ejecute el archivo fulcrum2.go mediante `go run fulcrum2.go`. En la otra, dirígase a la carpeta **"Thrawn"** y ejecute el archivo Thrawn.go mediante `go run Thrawn.go`.
- En las dos siguientes, conéctese a la máquina **"dist31"**. En una de ellas, dirígase a la carpeta **"Fulcrum3"** y ejecute el archivo fulcrum3.go mediante `go run fulcrum3.go`. En la otra, dirígase a la carpeta **"Leia"** y ejecute el archivo leia.go mediante `go run leia.go`.
- En la última terminal, debe conectarse a la máquina **"dist32"**. En ella, dirígase a la carpeta **"Broker"** y ejecute el archivo broker.go mediante `go run broker.go`
- **Nota:** En cada una de las carpetas, se encontrará otro readme donde se detalla qué función cumple cada uno de los programas y cómo lo hace.

#### Instrucciones de Uso:
- Una vez que todas las terminales se encuentren conectadas, puede comenzar a utilizar aquellas con menú interactivo, estas son: Ahsoka (en dist29), Thrawn (en dist30) y Leia (en dist31)
- Tanto en **Ahsoka** como en **Thrawn** puede realizar las siguientes acciones siguiendo las indicaciones mostradas por pantalla: Añadir una ciudad (a un planeta existente o uno nuevo a crear), cambiar el nombre de una ciudad en un planeta, cambiar la cantidad de rebeldes y eliminar alguna ciudad. Estas consultas son enviadas al **Servidor Broker**, donde aleatoriamente seleccionará entre los 3 **fulcrum** y devolverá la dirección de ellos. A continuación, Ahsoka o Thrawn se conectará al fulcrum, donde en este último se crearán, editarán o borrarán los documentos correspondientes.
- Por otro lado, en **Leia** se pueden consultar los rebeldes existentes en tal ciudad. Esta también se conectará al **Broker**, quien aleatoriamente seleccionará un **fulcrum** y consultará en aquél fulcrum por la ciudad pedida. Cabe destacar que cuando no existe la ciudad en ese Fulcrum (porque no se realizó el merge), indicará que la cantidad de rebeldes es cero. 