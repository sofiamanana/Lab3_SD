package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)





func main(){
	//Comienza conexion con el broker
	log.Printf("Informante Ahsoka Tano iniciada. \n")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.172:9050", grpc.WithInsecure()) //Conexión a Broker
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewInformanteBrokerClient(conn)

	log.Printf("¿Qué acción desea realizar?:\n")
	log.Printf("[1] Añadir una nueva ciudad.\n")
	log.Printf("[2] Actualizar nombre ciudad.\n")
	log.Printf("[3] Actualizar número de rebeldes de ciudad.\n")
	log.Printf("[4] Borrar una ciudad.\n")
	log.Printf("[5] No quiero hacer ni una wea más.\n")
	var opcion int32 = 0
	for ok:= true; ok; ok = (opcion!=5){
		fmt.Scan(&opcion)
		if (opcion==1){
			response, err := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "AddCity"})
			if err != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err)
			}
			log.Printf("Respuesta del Broker: %d", response.Valor)
			log.Printf("¿Cuál es el nombre del planeta?:\n")
			var planeta string
			fmt.Scan(&planeta)
			log.Printf("¿Cuál es el nombre de la ciudad?:\n")
			var ciudad string
			fmt.Scan(&ciudad)
			log.Printf("¿Cuántos rebeldes?:\n")
			var rebeldes int32
			fmt.Scan(&rebeldes)
			//aquí debería enviar los datos al fulcrum de response.valor 
		}
	}
	
	
}




