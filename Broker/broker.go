package main

import (
	pb "Lab3_SD/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Server struct {
	pb.UnimplementedBrokerServer
}

func (s *Server) GetNumberRebels(ctx context.Context, in *pb.PlanetaCiudad) (*pb.Numero, error) {
	split := strings.Split(in.Body, ",")
	planeta := split[0]
	ciudad := split[1]
	log.Printf("Leia pregunto por el planeta %s y la ciudad %s", planeta, ciudad)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.171:9040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewFulcrumClient(conn)
	response, err := c.PreguntarInformantes(context.Background(), &pb.PlanetaCiudad{Body: in.Body})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	return &pb.Numero{Num: response.Num}, nil
}

func (inf *Server) QuieroHacer(ctx context.Context, in *pb.Comando) (*pb.Redirigido, error) {
	var fulcrum int32 = 0
	var ip string
	log.Printf("El informante desea hacer: %s", in.Comando)
	fulcrum = rand.Int31n(3) + 1 //Se escoge un numero al azar entre 1 y 3 (corresponden a los 3 fulcrum)
	if fulcrum == 1 {            //ip del dist29 10.6.40.169
		ip = "10.6.40.169"
		return &pb.Redirigido{Valor: ip}, nil
	} else if fulcrum == 2 { //ip del dist30 10.6.40.170
		ip = "10.6.40.170"
		return &pb.Redirigido{Valor: ip}, nil
	} else { //ip del dist31 10.6.40.171
		ip = "10.6.40.171"
		return &pb.Redirigido{Valor: ip}, nil
	}
}

func ServerInformante() { //Conexión para conectar este broker (servidor) al informante (cliente)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	inf := grpc.NewServer()

	pb.RegisterBrokerServer(inf, &Server{})
	if err := inf.Serve(lis); err != nil {
		log.Fatalf("falló la conexión informante-broker: %s", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) //para seleccionar el server fulcrum al azar
	ServerInformante()
}
