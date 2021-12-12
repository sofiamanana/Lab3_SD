package main

import (
	"context"
	"log"

	pb "Lab3_SD/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.172:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewBrokerClient(conn)
	response, err := c.GetNumberRebels(context.Background(), &pb.PlanetaCiudad{Body: "Chile,Santiago"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Respuesta del Broker: %s", response.Num)
}
