## Laboratorio 3 - Sistemas Distribuidos

#### Informante Ahsoka Tano:
- Se conecta al Servidor Broker, enviándole el comando a realizar, recibe como respuesta la dirección del fulcrum al cual dirigirse.
- Se conecta a un Fulcrum aleatorio, enviándole el mismo comando que al Broker. Recibe como respuesta del fulcrum el reloj de vector del registro planetario del planeta que se modificó.
- En el caso de agregar ciudad, si se quiere crear una ciudad sin rebeldes, ingresar 0 como respuesta cuando se pregunte por los rebeldes.
- Read your Writes es implementado dentro de las funciones del servicio, con la función LeerArchivo que lee el archivo recién escrito y comprueba que sea la última versión dentro del Fulcrum que la invocó.
