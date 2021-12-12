package main

import (
	"context"
	"log"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.172:9050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewBrokerClient(conn)
	log.Printf("Leia Organa iniciada")
	log.Printf("¿Qué acción desea realizar?:\n")
	log.Printf("[1] Preguntar informacion.\n")
	var opcion int32 = 0
	for ok := true; ok; ok = (opcion != 5) {
		fmt.Scan(&opcion)
		if opcion == 1 {
			log.Printf("¿Planeta?\n")
			fmt.Scan(&planeta)
			log.Printf("¿Ciudad?\n")
			fmt.Scan(&ciudad)
			str := []string{planeta,ciudad}
			res := strings.Join(str, ",")
			response, err := c.GetNumberRebels(context.Background(), &pb.PlanetaCiudad{Body: res})
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			log.Printf("Respuesta del Broker: %s", response.Num)
			}
		}
	}


	
}
