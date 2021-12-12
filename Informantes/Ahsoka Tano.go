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

func main() {
	//--------------CONEXIONES FULCRUM ------------
	var conn1 *grpc.ClientConn //FULCRUM 1
	conn1, err1 := grpc.Dial("localhost:9060", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}
	defer conn1.Close()
	fulcrum1 := pb.NewFulcrumClient(conn1)

	var conn2 *grpc.ClientConn //FULCRUM 2
	conn2, err2 := grpc.Dial("10.6.40.170:9070", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}
	defer conn2.Close()
	fulcrum2 := pb.NewFulcrumClient(conn2)

	var conn3 *grpc.ClientConn //FULCRUM 3
	conn3, err3 := grpc.Dial("10.6.40.171:9040", grpc.WithInsecure())
	if err3 != nil {
		log.Fatalf("did not connect: %s", err3)
	}
	defer conn3.Close()
	fulcrum3 := pb.NewFulcrumClient(conn3)
	// -------- FIN CONEXIONES FULCRUM ----------------

	//Comienza conexion con el broker
	log.Printf("Informante Ahsoka Tano iniciada. \n")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.172:9050", grpc.WithInsecure()) //Conexión a Broker
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewBrokerClient(conn)

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
			var planet, city string
			var rebelds int32
			planet, city, rebelds = AddCiudad()
			if response.Valor == "10.6.40.169" { //fulcrum1 localhots
				res_fulcrum1, err_f1 := fulcrum1.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f1 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f1)
				}
				log.Printf("Respuesta del Fulcrum 1: %d", res_fulcrum1.X)

			} else if response.Valor == "10.6.40.170" { //fulcrum2
				res_fulcrum2, err_f2 := fulcrum2.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f2 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f2)
				}
				log.Printf("Respuesta del Fulcrum 2: %d", res_fulcrum2.X)
			} else { //fulcrum 3
				res_fulcrum3, err_f3 := fulcrum3.AddCity(context.Background(), &pb.Estructura{Planeta: planet, Ciudad: city, Rebeldes: rebelds})
				if err_f3 != nil {
					log.Fatalf("Error when calling AddCity: %s", err_f3)
				}
				log.Printf("Respuesta del Fulcrum 3: %d", res_fulcrum3.X)

			}
		}
	}

}
