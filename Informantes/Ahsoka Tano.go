package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)





func main(){
	//Comienza conexion con el broker
	log.Printf("Informante Ahsoka Tano iniciada. \n")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.169:9000", grpc.WithInsecure()) //Conexión a Broker
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	c := pb.NewInformanteBrokerClient(conn)

	log.Printf("¿Qué acción desea realizar?:\n")
	log.Printf("[1] Añadir una nueva ciudad.\n")
	log.Printf("[2] Actualizar nombre ciudad.\n")
	log.Printf("[3] Actualizar número de rebeldes de ciudad.\n")
	log.Printf("[4] Borrar una ciudad.\n")
	log.Printf("[5] No quiero hacer ni una wea más.\n")
	fmt.Scan(&opcion)
	if (opcion==1){
		response, err := c.QuieroHacer(context.Background(), &pb.Comando{comando: "AddCity"})
		if err != nil {
		log.Fatalf("Error when calling QuieroHacer: %s", err)
		}
		log.Printf("Respuesta del Lider: %s", response.valor)
		log.Printf("¿Cuál es el nombre del planeta?:\n")
		fmt.Scan(&planeta)
		log.Printf("¿Cuál es el nombre de la ciudad?:\n")
		fmt.Scan(&ciudad)
		log.Printf("¿Cuántos rebeldes?:\n")
		fmt.Scan(&rebeldes)
		
	}
	defer conn.Close()	 //esta conexión debería cerrarse cuando el informante no quiera hacer nada más?
}


