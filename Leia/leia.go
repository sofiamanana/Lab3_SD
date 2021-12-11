package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.169:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewNumberRebelsClient(conn)
	response, err := c.GetNumberRebels(context.Background(), &pb.PlanetaCiudad{Body: "Chile,Santiago"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Respuesta del Broker: %s", response.Num)
}