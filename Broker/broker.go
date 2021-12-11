package main

import (
	"context"
	"fmt"
	"log"
	"rand"
	"net"
	"strings"
	"time"
	"google.golang.org/grpc"
	pb "Lab3_SD/proto"
)

type Server struct {
	pb.UnimplementedNumberRebelsServer
}

type Server1 struct {
	pb.UnimplementedInformanteBrokerServer
}

func (s *Server) GetNumberRebels(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Numero, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	log.Printf("Leia pregunto por el planeta %s y la ciudad %s",planeta, ciudad)
	return &pb.Numero{Num: 11}, nil
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

func (s *Server1) QuieroHacer(ctx context.Context, in *pb.Comando) (*pb.Redirigido, error) { 
	var fulcrum int32 = 0
	log.Printf("El informante desea hacer: %s", in.Comando)
	fulcrum = rand.Int31n(3) +1 //Se escoge un numero al azar entre 1 y 3 (corresponden a los 3 fulcrum)
	return &pb.Redirigido{Valor: fulcrum}, nil
}

func ServerInformante(){ //Conexión para conectar este broker (servidor) al informante (cliente)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterInformanteBrokerServer(s, &Server1{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-broker: %s", err)
	}
}

func main(){
	rand.Seed(time.Now().UnixNano())  //para seleccionar el server fulcrum al azar

}