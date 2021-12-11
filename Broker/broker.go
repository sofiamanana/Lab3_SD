package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)

func (s *Server) GetNumberRebels(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Numero, error) {
	split := strings.Split(in.PlanetaCiudad, ",")
	planeta := split[0]
	ciudad := split[1]
	log.Printf("Leia pregunto por el planeta %s y la ciudad %s",planeta, ciudad)
	return &pb.Numero{num: 11}, nil
}

func ServerLeia(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNumberRebelsServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func (s *Server) QuieroHacer(ctx context.Context, in *pb.comando) (*pb.valor, error) { 
	var fulcrum int32 = 0
	log.Printf("El informante desea hacer: %s", in.comando)
	fulcrum = rand.Int31n(3) +1 //Se escoge un numero al azar entre 1 y 3 (corresponden a los 3 fulcrum)
	return &pb.Redirigido{valor: fulcrum}, nil
}

func ServerInformante(){ //Conexión para conectar este broker (servidor) al informante (cliente)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterInformanteBrokerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-broker: %s", err)
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())  //para seleccionar el server fulcrum al azar

}