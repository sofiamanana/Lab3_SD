package main

import (
	"context"
	"fmt"
	"log"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

func AddCiudad() (planeta string, ciudad string, rebeldes int32) {
	log.Printf("¿Cuál es el nombre del planeta?:\n")
	fmt.Scan(&planeta)
	log.Printf("¿Cuál es el nombre de la ciudad?:\n")
	fmt.Scan(&ciudad)
	log.Printf("¿Cuántos rebeldes?:\n")
	fmt.Scan(&rebeldes)
	return
}

func ConexionFulcrum1(){ //esta es local
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial("localhost:9060", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer conn1.Close()
	fulcrum1 := pb.NewInformanteBrokerClient(conn1)
}


func ConexionFulcrum2(){ //esta es a dist30
	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial("10.6.40.170:9070", grpc.WithInsecure()) 
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer conn2.Close()
	fulcrum2 := pb.NewInformanteBrokerClient(conn2)
}

func ConexionFulcrum3(){ //esta es a dist31
	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial("10.6.40.171:9080", grpc.WithInsecure()) 
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer conn3.Close()
	fulcrum3 := pb.NewInformanteBrokerClient(conn3)
}



func main() {
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
	for ok := true; ok; ok = (opcion != 5) {
		fmt.Scan(&opcion)
		if opcion == 1 {
			response, err := c.QuieroHacer(context.Background(), &pb.Comando{Comando: "AddCity"})
			if err != nil {
				log.Fatalf("Error when calling QuieroHacer: %s", err)
			}
			log.Printf("Respuesta del Broker: %s", response.Valor)
			AddCiudad()
			//aquí debería enviar los datos al fulcrum de response.valor
		}
	}

}
