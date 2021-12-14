package main

import (
	pb "Lab3_SD/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strings"
)

var Vector = make(map[string][]int32)

func main() {
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

	var opcion int32 = 0
	var ciudad, planeta, ip, rebeldes string
	for ok := true; ok; ok = (opcion != 2) {
		log.Printf("Leia Organa iniciada")
		log.Printf("¿Qué acción desea realizar?:\n")
		log.Printf("[1] Preguntar informacion.\n")
		log.Printf("[2] No quiero hacer ni una wea mas.\n")
		fmt.Scan(&opcion)
		if opcion == 1 {
			log.Printf("¿Planeta?\n")
			fmt.Scan(&planeta)
			log.Printf("¿Ciudad?\n")
			fmt.Scan(&ciudad)
			Vector[planeta] = []int32{0, 0, 0, 0}
			str := []string{planeta, ciudad}
			res := strings.Join(str, ",")

			//Monotonic reads

			if Vector[planeta] != {0,0,0,0} {

				response, err := c.GetNumberRebels(context.Background(), &pb.PlanetaCiudad{Body: res})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				ip = response.Ip
				if ip == "10.6.40.169" {
					if Vector[planeta][1] < response.X {
						rebeldes = response.Body
					}
				} else if ip == "10.6.40.170" {
					if Vector[planeta][2] < response.X {
						rebeldes = response.Body
					}
				} else {
					if Vector[planeta][3] < response.X {
						rebeldes = response.Body
					}
				}

				log.Printf("En el planeta %s hay %s rebeldes", planeta, response.Body)

			} else {
				response, err := c.GetNumberRebels(context.Background(), &pb.PlanetaCiudad{Body: res})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}
				log.Printf("En el planeta %s hay %s rebeldes", planeta, response.Body)
				ip = response.Ip
				rebeldes = response.Body
				if ip == "10.6.40.169" {
					Vector[planeta][0] = 1
					Vector[planeta][1] = response.X
					Vector[planeta][2] = response.Y
					Vector[planeta][3] = response.Z
				} else if ip == "10.6.40.170" {
					Vector[planeta][0] = 2
					Vector[planeta][1] = response.X
					Vector[planeta][2] = response.Y
					Vector[planeta][3] = response.Z
				} else {
					Vector[planeta][0] = 3
					Vector[planeta][1] = response.X
					Vector[planeta][2] = response.Y
					Vector[planeta][3] = response.Z
				}
			}

		}
	}
}
